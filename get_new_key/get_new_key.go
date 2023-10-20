package get_new_key

import (
	"net/http"

	"github.com/Flajt/decentproof-backend/helper"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func HandleGetNewKey(w http.ResponseWriter, r *http.Request) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	log.Info().Msg("Get new key request")
	authHeader := r.Header.Get("X-Appcheck")
	if authHeader == "" {
		log.Error().Msg("No auth header")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized"))
		return
	}
	appCheckWrapper := NewAppcheckWrapper()
	success, err := appCheckWrapper.CheckApp(authHeader)
	if err != nil {
		log.Error().Msg("Invalid App Check token")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized"))
		return
	} else {
		log.Info().Msg("No Error validating AppCheck token")
		apiKeys := helper.RetrievApiKeys()
		if success {
			log.Info().Msg("Valid App Check token,adding api key to response")
			w.WriteHeader(http.StatusOK)
			if len(apiKeys) == 1 {
				w.Write([]byte(apiKeys[0]))
			} else if len(apiKeys) == 2 {
				w.Write([]byte(apiKeys[1]))
			} else {
				log.Debug().Msgf("%d api keys found", len(apiKeys))
				panic("Invalid number of api keys found!")
			}
			return
		} else {
			log.Error().Msg("Invalid App Check token")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
			return
		}
	}
}
