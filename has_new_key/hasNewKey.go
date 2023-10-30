package has_new_key

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/Flajt/decentproof-backend/helper"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func HandleHasNewKey(w http.ResponseWriter, r *http.Request) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	log.Info().Msg("Has new key request")
	keys := helper.RetrievApiKeys()
	requestKey := strings.Split(r.Header.Get("Authorization"), " ")[1]
	isValid := helper.VerifyApiKey(r, keys)
	if isValid {
		log.Info().Msg("Valid API key, checking if new key is available")
		if keys[0] == requestKey {
			log.Trace().Msg("key[0] == requestKey")
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
		} else if len(keys) > 1 && keys[1] == requestKey { // in case we don't have two keys to prevent crashes
			log.Trace().Msg("key[1] == requestKey")
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
		} else {
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
