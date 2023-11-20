package decentproof_cronjob

import (
	"encoding/base64"
	"net/http"
	"os"
	"sort"
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
		returnError(w)
		panic(err)
	}
	if secretHolder.TotalCount == 0 {
		log.Debug().Msg("No secrets found, creating new one")
		apiKey := helper.GenerateApiKey(32)
		keyAsBytes := []byte(apiKey)
		if _, err := wrapper.SetSecret("apiKey", keyAsBytes); err != nil {
			log.Fatal().Msg(err.Error())
			returnError(w)
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
						returnError(w)
						panic(err)
					}
				} else {
					log.Info().Msg("Second secret is newer, deleting and replacing first secret")

					if err := wrapper.DeleteSecretVersion(firstSecret.SecretID, strconv.FormatUint(uint64(firstSecret.Revision), 10)); err != nil {
						log.Fatal().Msg(err.Error())
						returnError(w)
						panic(err)
					}
					apiKey := helper.GenerateApiKey(32)
					apiKeyBytes := []byte(apiKey)
					//TODO: Can this duplication be removed?
					if firstSecretCreationDateLater {
						if err := wrapper.CreateNewSecretVersion(*secretHolder.Secrets[1], apiKeyBytes); err != nil {
							log.Fatal().Msg(err.Error())
							returnError(w)
							panic(err)
						}
					} else {
						if err := wrapper.CreateNewSecretVersion(*secretHolder.Secrets[0], apiKeyBytes); err != nil {
							log.Fatal().Msg(err.Error())
							returnError(w)
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
					returnError(w)
					panic(err)
				}
			}
			// Handle E-Mail encryption key management
			secretHolder, err := wrapper.ListSecrets("ENCRYPTION_KEY")
			if err != nil {
				log.Fatal().Msg(err.Error())
				returnError(w)
			}
			if secretHolder.TotalCount == 0 {
				base64Key := helper.GenerateApiKey(32)
				bytes, err := base64.StdEncoding.DecodeString(base64Key)
				if err != nil {
					log.Fatal().Msg(err.Error())
					returnError(w)
				}
				wrapper.SetSecret("ENCRYPTION_KEY", bytes)
			} else if secretHolder.TotalCount > 1 {
				log.Fatal().Msg("More than one encryption key found, this should not happen")
				returnError(w)
			} else {
				log.Info().Msg("Encryption key found!")
				secretVersion, err := wrapper.ListSecretVersions(secretHolder.Secrets[0].ID)
				if err != nil {
					log.Fatal().Msg(err.Error())
					returnError(w)
				}
				if secretVersion.TotalCount == 0 {
					log.Info().Msg("No secret versions found, creating new one")
					base64Key := helper.GenerateApiKey(32)
					bytes, err := base64.StdEncoding.DecodeString(base64Key)
					if err != nil {
						log.Fatal().Msg(err.Error())
						returnError(w)
					}
					wrapper.CreateNewSecretVersion(*secretHolder.Secrets[0], bytes)
				} else if secretVersion.TotalCount == 1 {
					log.Info().Msg("One secret version found, creating new one")
					base64Key := helper.GenerateApiKey(32)
					bytes, err := base64.StdEncoding.DecodeString(base64Key)
					if err != nil {
						log.Fatal().Msg(err.Error())
						returnError(w)
					}
					wrapper.CreateNewSecretVersion(*secretHolder.Secrets[0], bytes)
				} else if secretVersion.TotalCount == 2 {
					sort.Slice(secretVersion.SecretVersions, func(i, j int) bool {
						return secretVersion.SecretVersions[i].CreatedAt.Before(*secretVersion.SecretVersions[j].CreatedAt)
					})
					log.Info().Msg("Two secret versions found, deleting oldest and replacing it")
					err = wrapper.DeleteSecretVersion(secretVersion.SecretVersions[0].SecretID, strconv.FormatUint(uint64(secretVersion.SecretVersions[1].Revision), 10))
					if err != nil {
						log.Fatal().Msg(err.Error())
						returnError(w)
					}
					base64Key := helper.GenerateApiKey(32)
					bytes, err := base64.StdEncoding.DecodeString(base64Key)
					if err != nil {
						log.Fatal().Msg(err.Error())
						returnError(w)
					}
					err = wrapper.CreateNewSecretVersion(*secretHolder.Secrets[0], bytes)
					if err != nil {
						log.Fatal().Msg(err.Error())
						returnError(w)
					}
				} else {
					log.Fatal().Msg("More than two secret versions found, this should not happen")
					returnError(w)
				}
			}

			log.Info().Msg("Done")
			w.Header().Set("Content-type", "text/plain")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Done"))
		}

	}
}

func returnError(w http.ResponseWriter) {
	w.Header().Set("Content-type", "text/plain")
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("Something went wrong"))
}
