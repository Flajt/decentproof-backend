package webhook

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	originstamp "github.com/Flajt/decentproof-backend/originstamp"
	models "github.com/Flajt/decentproof-backend/originstamp/models"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	mail "github.com/xhit/go-simple-mail/v2"
)

func HandleWebhookCallBack(w http.ResponseWriter, r *http.Request) {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Info().Msg("Webhook callback request")
	encryptedMailAddress := r.URL.Query().Get("mail")
	if encryptedMailAddress == "" || encryptedMailAddress == " " {
		log.Info().Msg("No email address provided")
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Done"))
		return
	}

	orignStampApi := originstamp.NewOriginStampApiClient(os.Getenv("ORIGINSTAMP_API_KEY"))
	var requestBody models.OriginStampWebhookRequestBody
	json.NewDecoder(r.Body).Decode(&requestBody)

	defer r.Body.Close()

	proofBody := models.OriginStampProofRequestBody{Currency: requestBody.Timestamps[0].CurrencyID, HashString: requestBody.HashString, ProofType: 1}
	data, err := orignStampApi.GetProof(proofBody)
	if err != nil {
		log.Error().Err(err).Msg("Error getting proof from OriginStamp API")
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
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
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
	fileBytes, err := io.ReadAll(fileRequest.Body)
	if err != nil {
		log.Error().Err(err).Msg("Error reading file from OriginStamp API")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Something went wrong"))
		return
	}
	defer fileRequest.Body.Close()

	log.Info().Msg("Setting up email response")
	//TODO: Consider moving this in a seperate function
	smtpServer := "smtp.tem.scw.cloud"
	smtpPort := 587
	userName := os.Getenv("SCW_DEFAULT_PROJECT_ID")
	password := os.Getenv("SCW_EMAIL_SECRET")

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

	attachment := mail.File{Data: fileBytes, Name: data.Data.FileName}
	finalizedMail := email.SetFrom(from).AddTo(encryptedMailAddress).SetSubject("Decentproof: Document Persisted").SetBody(mail.TextHTML, "This is an autogenerated E-Mail. Attached you will find your verification certificate.").Attach(&attachment)
	err = finalizedMail.Send(client)
	if err != nil {
		log.Error().Err(err).Msg("Error sending email")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Something went wrong"))
		return
	}
	log.Info().Msg("Successfully sent email")
	w.WriteHeader(http.StatusOK)
}
