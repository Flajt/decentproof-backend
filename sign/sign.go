package sign

import (
	"encoding/hex"
	"encoding/json"
	"net/http"
	"os"

	encryption_service "github.com/Flajt/decentproof-backend/encryption"
	"github.com/Flajt/decentproof-backend/helper"
	"github.com/Flajt/decentproof-backend/originstamp"
	models "github.com/Flajt/decentproof-backend/originstamp/models"

	secret_wrapper "github.com/Flajt/decentproof-backend/scw_secret_wrapper"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func HandleSignature(w http.ResponseWriter, r *http.Request) {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Info().Msg("Signature request received")
	scwWrapper := secret_wrapper.NewScaleWayWrapperFromEnv()
	isValid := helper.VerifyApiKey(r, helper.RetrievApiKeys(scwWrapper))
	if !isValid {
		log.Error().Msg("Unauthorized request")
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized"))
		return
	}

	scw_wrapper := secret_wrapper.NewScaleWayWrapperFromEnv()
	signatureManager := NewSignatureManager(scw_wrapper)
	signatureManager.InitSignatureManager()
	jsonDecoder := json.NewDecoder(r.Body)
	var holder SignatureRequestBody
	if err := jsonDecoder.Decode(&holder); err != nil {
		log.Error().Msg("Bad Request, can't decode json body. Details: " + err.Error())
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request"))
		return
	}
	signatureBytes, err := signatureManager.SignData([]byte(holder.Data))
	if err != nil {
		log.Error().Msg("Internal Server Error, can't sign data. Details: " + err.Error())
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}
	APIKEY := os.Getenv("ORIGINSTAMP_API_KEY")
	webhookUrl := os.Getenv("WEBHOOK_URL")
	if holder.Email != "" {
		encryptionService := encryption_service.NewEncryptionService(scw_wrapper)
		encryptionData, err := encryptionService.EncryptData([]byte(holder.Email))
		if err != nil {
			log.Error().Err(err).Msg("Can't encrypt email")
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		webhookUrl += "?mail=" + hex.EncodeToString(encryptionData.Data) + "&nonce=" + hex.EncodeToString(encryptionData.Nonce)
	} else {
		webhookUrl = ""
	}
	client := originstamp.NewOriginStampApiClient(APIKEY)
	bitcoinNotificationTarget := models.OriginStampNotificationTarget{Target: webhookUrl, NotificationType: 1, Currency: 0}
	etheriumNotificationTarget := models.OriginStampNotificationTarget{Target: webhookUrl, NotificationType: 1, Currency: 1}
	timeStampReqModel := models.OriginStampTimestampRequestBody{Comment: hex.EncodeToString(signatureBytes), Hash: holder.Data, Notifications: []models.OriginStampNotificationTarget{bitcoinNotificationTarget, etheriumNotificationTarget}}
	resp, err := client.CreateTimestamp(timeStampReqModel)
	if err != nil {
		log.Error().Err(err).Msg("Can't create timestamp. Details: " + err.Error())
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if resp.ErrorMessage != "" {
		log.Error().Msg("Can't create timestamp. Details: " + resp.ErrorMessage)
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	log.Info().Msg("Signature request completed")
}
