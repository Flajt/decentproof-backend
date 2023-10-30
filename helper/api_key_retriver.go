package helper

import (
	"os"
	"strconv"

	scw_secret_manager "github.com/Flajt/decentproof-backend/scw_secret_wrapper"
)

func RetrievApiKeys() []string {
	scwSetupData := scw_secret_manager.ScaleWaySetupData{}
	scwSetupData.AccessKey = os.Getenv("SCW_ACCESS_KEY")
	scwSetupData.ProjectID = os.Getenv("SCW_DEFAULT_PROJECT_ID")
	scwSetupData.SecretKey = os.Getenv("SCW_SECRET_KEY")
	scwSetupData.Region = os.Getenv("SCW_DEFAULT_REGION")
	client := scw_secret_manager.NewScaleWayWrapper(scwSetupData)
	secrets, err := client.ListSecrets("apiKey")
	if err != nil {
		panic(err)
	}
	if len(secrets.Secrets) == 0 { // Would be fundamentally broken if this happened
		panic("No api keys found")
	} else {
		secretVersions, err := client.ListSecretVersions(secrets.Secrets[0].ID)
		if err != nil { // Should be impossible if there is at least one secret,since it's created with it
			panic(err)
		}
		var apiKeys []string
		for _, secret := range secretVersions.SecretVersions {
			if secret.Status != "destroyed" {
				data, err := client.GetSecretData("apiKey", strconv.FormatUint(uint64(secret.Revision), 10))
				apiKeys = append(apiKeys, string(data))
				if err != nil {
					/// Should also be impossible to not get the data if the rest is true
					panic(err)
				}
			}
		}
		return apiKeys
	}
}
