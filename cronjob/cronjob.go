package decentproof_cronjob

import (
	"net/http"
	"os"
	"strconv"

	scw_secret_manager "github.com/Flajt/decentproof-backend/scw_secret_wrapper"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	var setupData = scw_secret_manager.ScaleWaySetupData{ProjectID: os.Getenv("SCW_DEFAULT_PROJECT_ID"), AccessKey: os.Getenv("SCW_ACCESS_KEY"), SecretKey: os.Getenv("SCW_SECRET_KEY"), Region: os.Getenv("SCW_DEFAULT_REGION")}

	//TODO; Find a way to make things less messy
	wrapper := scw_secret_manager.NewScaleWayWrapper(setupData)

	// Step one: Check for existing keys
	secretHolder, err := wrapper.ListSecrets("apiKey")
	if err != nil {
		returnError(w, err)
		panic(err)
	}
	if secretHolder.TotalCount == 0 {
		apiKey := GenerateApiKey()
		keyAsBytes := []byte(apiKey)
		if _, err := wrapper.SetSecret("apiKey", keyAsBytes); err != nil {
			panic(err)
		}
		w.Header().Set("Content-type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Done"))
	} else {
		if versionHolder, err := wrapper.ListSecretVersions(secretHolder.Secrets[0].ID); err != nil {
			panic(err)
		} else {
			if versionHolder.TotalCount == 2 {
				firstSecret := versionHolder.SecretVersions[0]
				secondSecret := versionHolder.SecretVersions[1]
				secretOneCreationTime := firstSecret.CreatedAt
				secretTwoCreationTime := secondSecret.CreatedAt
				firstSecretCreationDateLater := secretOneCreationTime.After(*secretTwoCreationTime)
				if firstSecretCreationDateLater {
					//Delete secret 2
					if err := wrapper.DeleteSecretVersion(secondSecret.SecretID, strconv.FormatUint(uint64(secondSecret.Revision), 10)); err != nil {
						returnError(w, err)
						panic(err)
					}
				} else {
					if err := wrapper.DeleteSecretVersion(firstSecret.SecretID, strconv.FormatUint(uint64(firstSecret.Revision), 10)); err != nil {
						returnError(w, err)
						panic(err)
					}
					apiKey := GenerateApiKey()
					apiKeyBytes := []byte(apiKey)
					//TODO: Can this duplication be removed?
					if firstSecretCreationDateLater {
						if err := wrapper.CreateNewSecretVersion(*secretHolder.Secrets[1], apiKeyBytes); err != nil {
							returnError(w, err)
							panic(err)
						}
					} else {
						if err := wrapper.CreateNewSecretVersion(*secretHolder.Secrets[0], apiKeyBytes); err != nil {
							returnError(w, err)
							panic(err)
						}
					}
				}
			} else {
				apiKey := GenerateApiKey()
				apiKeyBytes := []byte(apiKey)
				if err := wrapper.CreateNewSecretVersion(*secretHolder.Secrets[0], apiKeyBytes); err != nil {
					returnError(w, err)
					panic(err)
				}
			}
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
