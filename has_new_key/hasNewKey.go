package has_new_key

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/Flajt/decentproof-backend/helper"
	scw_secret_manager "github.com/Flajt/decentproof-backend/scw_secret_wrapper"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// TODO: See if that can be written better
func HandleHasNewKey(w http.ResponseWriter, r *http.Request) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	log.Info().Msg("Has new key request")
	scw_secret_manager := scw_secret_manager.NewScaleWayWrapperFromEnv()
	keys := helper.RetrievApiKeys(scw_secret_manager)

	requestKey := strings.Split(r.Header.Get("Authorization"), " ")[1]
	isValid := helper.VerifyApiKey(r, keys)
	if isValid {
		log.Info().Msg("Valid API key, checking if new key is available")
		if keys[0] == requestKey && len(keys) > 1 { // We have more than one key available, so this one should be the oldest
			log.Trace().Msg("key[0] == requestKey && len > 1")
			response := map[string]bool{"hasNewKey": true}
			responseBytes, err := json.Marshal(response)
			if err != nil {
				log.Error().Msg("Error marshalling response. Details: " + err.Error())
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("Internal Server Error"))
				return
			}
			log.Info().Msg("Successfully handled request")
			log.Trace().Msg("hasNewKey:true")
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(responseBytes)
			return
		} else if keys[0] == requestKey && len(keys) == 1 { // We only have one key available and it's the one we got
			log.Trace().Msg("key[0] == requestKey && len == 1")
			response := map[string]bool{"hasNewKey": false}
			responseBytes, err := json.Marshal(response)
			if err != nil {
				log.Error().Msg("Error marshalling response. Details: " + err.Error())
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("Internal Server Error"))
				return
			}
			log.Info().Msg("Successfully handled request")
			log.Trace().Msg("hasNewKey:false")
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(responseBytes)
			return

		} else if len(keys) > 1 && keys[1] == requestKey { // in case we don't have two keys to prevent crashes
			log.Trace().Msg("key[1] == requestKey && len > 1")
			response := map[string]bool{"hasNewKey": false}
			responseBytes, err := json.Marshal(response)
			if err != nil {
				log.Error().Msg("Error marshalling response. Details: " + err.Error())
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("Internal Server Error"))
				return
			}
			log.Info().Msg("Successfully handled request")
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(responseBytes)
			return
		} else { // everything else
			log.Trace().Msg("key[0] != requestKey && key[1] != requestKey")
			response := map[string]bool{"hasNewKey": true}
			responseBytes, err := json.Marshal(response)
			if err != nil {
				log.Error().Msg("Error marshalling response. Details: " + err.Error())
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("Internal Server Error"))
				return
			}
			log.Info().Msg("Successfully handled request")
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(responseBytes)
			return
		}

	} else {
		log.Trace().Msg("Invalid API key: " + requestKey)
		log.Info().Msg("Invalid API key, return hasNewKey: true")
		w.WriteHeader(http.StatusOK)
		response := map[string]bool{"hasNewKey": true}
		responseBytes, err := json.Marshal(response)
		if err != nil {
			log.Error().Msg("Error marshalling response. Details: " + err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal Server Error"))
			return
		}
		log.Info().Msg("Successfully handled request")
		w.Header().Set("Content-Type", "application/json")
		w.Write(responseBytes)
		return
	}
}
