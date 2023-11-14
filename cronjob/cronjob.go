package decentproof_cronjob

import (
	"net/http"
	"os"
	"strconv"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/Flajt/decentproof-backend/helper"
	scw_secret_manager "github.com/Flajt/decentproof-backend/scw_secret_wrapper"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Info().Msg("Starting cronjob")
	var setupData = scw_secret_manager.ScaleWaySetupData{ProjectID: os.Getenv("SCW_DEFAULT_PROJECT_ID"), AccessKey: os.Getenv("SCW_ACCESS_KEY"), SecretKey: os.Getenv("SCW_SECRET_KEY"), Region: os.Getenv("SCW_DEFAULT_REGION")}

	//TODO; Find a way to make things less messy
	wrapper := scw_secret_manager.NewScaleWayWrapper(setupData)

	// Step one: Check for existing keys
	secretHolder, err := wrapper.ListSecrets("apiKey")
	if err != nil {
		log.Error().Msg(err.Error())
		returnError(w, err)
		panic(err)
	}
	if secretHolder.TotalCount == 0 {
		log.Debug().Msg("No secrets found, creating new one")
		apiKey := helper.GenerateApiKey(32)
		keyAsBytes := []byte(apiKey)
		if _, err := wrapper.SetSecret("apiKey", keyAsBytes); err != nil {
			log.Fatal().Msg(err.Error())
			returnError(w, err)
			panic(err)
		}
		w.Header().Set("Content-type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Done"))
	} else {
		if versionHolder, err := wrapper.ListSecretVersions(secretHolder.Secrets[0].ID); err != nil {
			log.Fatal().Msg(err.Error())
			panic(err)
		} else {
			if versionHolder.TotalCount == 2 {
				log.Info().Msg("Two secrets found, deleting and relacing oldest")
				firstSecret := versionHolder.SecretVersions[0]
				secondSecret := versionHolder.SecretVersions[1]
				secretOneCreationTime := firstSecret.CreatedAt
				secretTwoCreationTime := secondSecret.CreatedAt
				firstSecretCreationDateLater := secretOneCreationTime.After(*secretTwoCreationTime)
				log.Info().Msg("Is first secret newer: " + strconv.FormatBool(firstSecretCreationDateLater))
				if firstSecretCreationDateLater {
					log.Debug().Msg("First secret is newer, deleting and replacing second secret")
					//Delete secret 2
					if err := wrapper.DeleteSecretVersion(secondSecret.SecretID, strconv.FormatUint(uint64(secondSecret.Revision), 10)); err != nil {
						log.Fatal().Msg(err.Error())
						returnError(w, err)
						panic(err)
					}
				} else {
					log.Info().Msg("Second secret is newer, deleting and replacing first secret")

					if err := wrapper.DeleteSecretVersion(firstSecret.SecretID, strconv.FormatUint(uint64(firstSecret.Revision), 10)); err != nil {
						log.Fatal().Msg(err.Error())
						returnError(w, err)
						panic(err)
					}
					apiKey := helper.GenerateApiKey(32)
					apiKeyBytes := []byte(apiKey)
					//TODO: Can this duplication be removed?
					if firstSecretCreationDateLater {
						if err := wrapper.CreateNewSecretVersion(*secretHolder.Secrets[1], apiKeyBytes); err != nil {
							log.Fatal().Msg(err.Error())
							returnError(w, err)
							panic(err)
						}
					} else {
						if err := wrapper.CreateNewSecretVersion(*secretHolder.Secrets[0], apiKeyBytes); err != nil {
							log.Fatal().Msg(err.Error())
							returnError(w, err)
							panic(err)
						}
					}
				}
			} else {
				log.Debug().Msg("One secret found, creating new one")
				apiKey := helper.GenerateApiKey(32)
				apiKeyBytes := []byte(apiKey)
				if err := wrapper.CreateNewSecretVersion(*secretHolder.Secrets[0], apiKeyBytes); err != nil {
					log.Fatal().Msg(err.Error())
					returnError(w, err)
					panic(err)
				}
			}
			log.Info().Msg("Done")
			w.Header().Set("Content-type", "text/plain")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Done"))
		}

	}
}

func returnError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-type", "text/plain")
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(err.Error()))
}
