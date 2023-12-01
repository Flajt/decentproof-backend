package verify_hash

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/Flajt/decentproof-backend/helper"
	"github.com/Flajt/decentproof-backend/originstamp"
	scw_secret_manager "github.com/Flajt/decentproof-backend/scw_secret_wrapper"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func HandleHashVerification(w http.ResponseWriter, r *http.Request) {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Info().Msg("Handling hash verification")
	var scwWrapper scw_secret_manager.IScaleWayWrapper
	if os.Getenv("DEBUG") == "TRUE" {
		log.Info().Msg("DEBUG MODE: TRUE")
		scwWrapper = scw_secret_manager.NewScaleWayWrapperForDev()
	} else {
		log.Info().Msg("DEBUG MODE: FALSE")
		scwWrapper = scw_secret_manager.NewScaleWayWrapperFromEnv()
	}
	apiKeys := helper.RetrievApiKeys(scwWrapper)
	isValid, err := helper.Authenticate(r, apiKeys, true)
	if !isValid || err != nil {
		log.Err(err).Msg("Invalid api key")
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized"))
		return
	}
	requestModel := VerifyRequestBody{}
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Error().Err(err).Msg("Error reading request body")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	defer r.Body.Close()
	err = json.Unmarshal(bytes, &requestModel)
	if err != nil {
		log.Error().Err(err).Msg("Error unmarshalling request body")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	if requestModel.Hash == "" {
		log.Error().Msg("A hash is required")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Hash is required"))
		return
	}
	ORIGINSTAMP_API_KEY := os.Getenv("ORIGINSTAMP_API_KEY")
	orignstampClient := originstamp.NewOriginStampApiClient(ORIGINSTAMP_API_KEY)
	resp, err := orignstampClient.GetTimestampStatus(requestModel.Hash)
	if err != nil {
		log.Error().Err(err).Msg("Error getting timestamp status")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	if resp.ErrorMessage != "" {
		log.Error().Msg(resp.ErrorMessage)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(resp.ErrorMessage))
		return
	}
	log.Info().Msg("Successfully verified hash")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)

}
