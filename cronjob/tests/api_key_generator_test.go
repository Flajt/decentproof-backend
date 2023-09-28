package decentproof_cronjob

import (
	"testing"

	decentproof_cronjob "github.com/Flajt/decentproof-backend/decentproof-cron"
)

func TestKeyGeneration(t *testing.T) {
	if key := decentproof_cronjob.GenerateApiKey(); key == "" {
		t.Error("Key is empty")
	} else {
		if len(key)%4 != 0 {
			t.Error("Key is not a multiple of 4")
			t.Log(len(key))
		}
	}

}
