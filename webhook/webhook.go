package webhook

import (
	"encoding/hex"
	"encoding/json"
	"io"
	"net/http"
	"os"

	encryption_service "github.com/Flajt/decentproof-backend/encryption"
	originstamp "github.com/Flajt/decentproof-backend/originstamp"
	models "github.com/Flajt/decentproof-backend/originstamp/models"
	scw_secret_manager "github.com/Flajt/decentproof-backend/scw_secret_wrapper"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	mail "github.com/xhit/go-simple-mail/v2"
)

func HandleWebhookCallBack(w http.ResponseWriter, r *http.Request) {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Info().Msg("Webhook callback request")
	encryptedMailAddress := r.URL.Query().Get("mail")
	nonce := r.URL.Query().Get("nonce")
	if encryptedMailAddress == "" || encryptedMailAddress == " " && nonce == "" || nonce == " " {
		log.Info().Msg("No email address provided")
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Done"))
		return
	}
	var scwWrapper scw_secret_manager.IScaleWayWrapper
	if os.Getenv("DEBUG") == "TRUE" {
		log.Info().Msg("DEBUG MODE: TRUE")
		scwWrapper = scw_secret_manager.NewScaleWayWrapperForDev()
	} else {
		log.Info().Msg("DEBUG MODE: FALSE")
		scwWrapper = scw_secret_manager.NewScaleWayWrapperFromEnv()
	}
	orignStampApi := originstamp.NewOriginStampApiClient(os.Getenv("ORIGINSTAMP_API_KEY"))
	var requestBody models.OriginStampWebhookRequestBody
	json.NewDecoder(r.Body).Decode(&requestBody)

	defer r.Body.Close()

	var generalProofBodies []models.OriginStampProofRequestBody = make([]models.OriginStampProofRequestBody, len(requestBody.Timestamps))
	if len(requestBody.Timestamps) > 0 && len(requestBody.Timestamps) < 3 {
		for i, timestamp := range requestBody.Timestamps {
			proofBody := models.OriginStampProofRequestBody{Currency: timestamp.CurrencyID, HashString: requestBody.HashString, ProofType: 1}
			generalProofBodies[i] = proofBody
		}
	} else {
		log.Error().Msg("Invalid number of timestamps")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request"))
		return
	}
	attachments := make([]*mail.File, len(generalProofBodies)) // If you don't use pointers, the attachement loop will only attach the last file twice
	for i, proofBody := range generalProofBodies {
		data, err := orignStampApi.GetProof(proofBody)
		if err != nil {
			log.Err(err).Msg("Error getting proof from OriginStamp API")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Something went wrong"))
			return
		}
		if data.ErrorMessage != "" {
			log.Error().Msg("Error returned while getting proof from OriginStamp API. Details: " + data.ErrorMessage)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Something went wrong"))
			return
		}
		downloadUrl := data.Data.DownloadURL
		request, err := http.NewRequest("GET", downloadUrl, nil)
		if err != nil {
			log.Error().Err(err).Msg("Error creating request to download file from OriginStamp API")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error requesting certificate"))
			return
		}
		request.Header["Accept"] = []string{"application/octet-stream"}
		fileRequest, err := http.DefaultClient.Do(request)
		if err != nil {
			log.Error().Err(err).Msg("Error downloading file from OriginStamp API")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Something went wrong"))
			return
		}
		defer fileRequest.Body.Close()
		fileBytes, err := io.ReadAll(fileRequest.Body)
		if err != nil {
			log.Error().Err(err).Msg("Error reading file from OriginStamp API")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Something went wrong"))
			return
		}
		attachments[i] = &mail.File{Data: fileBytes, Name: data.Data.FileName, MimeType: "application/pdf"}
	}

	decryptionService := encryption_service.NewEncryptionService(scwWrapper)
	emailBytes, err := hex.DecodeString(encryptedMailAddress)
	if err != nil {
		log.Error().Err(err).Msg("Error decoding email")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Something went wrong"))
	}
	nonceBytes, err := hex.DecodeString(nonce)
	if err != nil {
		log.Error().Err(err).Msg("Error decoding nonce")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Something went wrong"))
	}
	decryptedMailAddressBytes, err := decryptionService.DecryptData(emailBytes, nonceBytes)
	if err != nil {
		log.Error().Err(err).Msg("Error decrypting email")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Something went wrong"))
	}
	emailAddress := string(decryptedMailAddressBytes)

	log.Info().Msg("Setting up email response")
	//TODO: Consider moving this in a seperate function
	smtpServer := "smtp.tem.scw.cloud"
	smtpPort := 587
	userName := os.Getenv("SCW_DEFAULT_PROJECT_ID")
	password := os.Getenv("EMAIL_SECRET")

	server := mail.NewSMTPClient()

	server.Host = smtpServer
	server.Port = smtpPort
	server.Username = userName
	server.Password = password
	server.Encryption = mail.EncryptionTLS
	from := "no-reply@decentproof.com"
	client, err := server.Connect()
	if err != nil {
		log.Error().Err(err).Msg("Error connecting to SMTP server")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Something went wrong"))
		return
	}

	email := mail.NewMSG()

	email.SetFrom(from).AddTo(emailAddress).SetSubject("Decentproof: Document Persisted").SetBody(mail.TextPlain, "This is an autogenerated E-Mail. Attached you will find your verification certificate/s.")
	for _, attachment := range attachments {
		email.Attach(attachment)
	}

	err = email.Send(client)
	if err != nil {
		log.Error().Err(err).Msg("Error sending email")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Something went wrong"))
		return
	}
	if email.Error != nil {
		log.Error().Err(email.Error).Msg("Error setting up email")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Something went wrong"))
		return
	}
	log.Info().Msg("Successfully sent email")
	w.WriteHeader(http.StatusOK)
}
