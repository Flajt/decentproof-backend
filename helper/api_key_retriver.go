package helper

import (
	"strconv"

	scw_secret_manager "github.com/Flajt/decentproof-backend/scw_secret_wrapper"
)

func RetrievApiKeys(scwWrapper scw_secret_manager.IScaleWayWrapper) []string {

	secrets, err := scwWrapper.ListSecrets("apiKey")
	if err != nil {
		panic(err)
	}
	if len(secrets.Secrets) == 0 { // Would be fundamentally broken if this happened
		panic("No api keys found")
	} else {
		secretVersions, err := scwWrapper.ListSecretVersions(secrets.Secrets[0].ID)
		if err != nil { // Should be impossible if there is at least one secret,since it's created with it
			panic(err)
		}
		var apiKeys []string
		for _, secret := range secretVersions.SecretVersions {
			if secret.Status != "destroyed" {
				data, err := scwWrapper.GetSecretData("apiKey", strconv.FormatUint(uint64(secret.Revision), 10))
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
