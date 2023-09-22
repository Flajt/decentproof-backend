package scw_secret_manager

import (
	"os"
	"testing"
)

// /Removes all created secrets, used for testing
func CleanUp(t *testing.T) {
	var setupData = ScaleWaySetupData{ProjectID: os.Getenv("SCW_DEFAULT_PROJECT_ID"), AccessKey: os.Getenv("SCW_ACCESS_KEY"), SecretKey: os.Getenv("SCW_SECRET_KEY"), Region: os.Getenv("SCW_DEFAULT_REGION")}

	wrapper := NewScaleWayWrapper(setupData)
	if secrets, err := wrapper.ListSecrets(); err != nil {
		t.Error(err)
		t.Log("MANUAL CLEAN UP REQUIRED !!!!!!")
	} else {
		for _, secret := range secrets.Secrets {
			if err := wrapper.DeleteSecret(secret.ID); err != nil {
				t.Error(err)
				t.Log("MANUAL CLEAN UP REQUIRED !!!!!!")
			}
		}

	}
}
