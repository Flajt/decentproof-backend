package scw_secret_manager

import (
	"testing"
)

// /Removes all created secrets, used for testing
func CleanUp(t *testing.T, IDs []string) {
	wrapper := NewScaleWayWrapperFromEnv()
	for _, id := range IDs {
		if err := wrapper.DeleteSecret(id); err != nil {
			t.Error(err)
			t.Log("MANUAL CLEAN UP REQUIRED !!!!!!")
		}
	}

}
