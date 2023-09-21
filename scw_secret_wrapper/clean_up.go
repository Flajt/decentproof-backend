package scw_secret_manager

import "testing"

// /Removes all created secrets, used for testing
func CleanUp(t *testing.T) {
	if wrapper, err := NewScaleWayWrapper(); err != nil {
		t.Error(err)
		t.Log("MANUAL CLEAN UP REQUIRED !!!!!!")
	} else {
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
}
