package sign

import (
	"encoding/json"
	"net/http"

	"github.com/Flajt/decentproof-backend/helper"
	secret_wrapper "github.com/Flajt/decentproof-backend/scw_secret_wrapper"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func HandleSignature(w http.ResponseWriter, r *http.Request) {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Info().Msg("Signature request received")
	isValid := helper.VerifyApiKey(r, helper.RetrievApiKeys())
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
	signatureResp := SignatureResponseBody{Signature: string(signatureBytes)}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	respBytes, err := json.Marshal(signatureResp)
	if err != nil {
		log.Error().Msg("Internal Server Error, can't marshal response. Details" + err.Error())
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}
	w.Write(respBytes)
	log.Info().Msg("Signature request completed")
}
