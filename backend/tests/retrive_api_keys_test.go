package decentproof_functions

import (
	"testing"

	helper "github.com/Flajt/decentproof-backend/decentproof-functions/helper"
	scw_secret_wrapper "github.com/Flajt/decentproof-backend/scw_secret_wrapper"
	"github.com/joho/godotenv"
)

func TestRetrievApiKeys(t *testing.T) {
	godotenv.Load("../.env")
	t.Run("with zero entries", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Log("Recovered from panic")

			}
		}()
		keys := helper.RetrievApiKeys()
		if len(keys) != 0 {
			t.Error("Got keys, wanted none")
		}
	})

	t.Run("with one secret entry", func(t *testing.T) {
		scw_wrapper := scw_secret_wrapper.NewScaleWayWrapperFromEnv()
		_, err := scw_wrapper.SetSecret("apiKey", []byte("test"))
		if err != nil {
			t.Error(err)
		}
		keys := helper.RetrievApiKeys()
		if len(keys) != 1 {
			t.Errorf("Got %d keys, wanted one", len(keys))
		}
		scw_secret_wrapper.CleanUp(t)
	})

	t.Run("with two secret entries", func(t *testing.T) {
		scw_wrapper := scw_secret_wrapper.NewScaleWayWrapperFromEnv()
		secret, err := scw_wrapper.SetSecret("apiKey", []byte("test"))
		if err != nil {
			t.Error(err)
		}
		err = scw_wrapper.CreateNewSecretVersion(*secret, []byte("test2"))
		if err != nil {
			t.Error(err)
		}
		keys := helper.RetrievApiKeys()
		if len(keys) != 2 {
			t.Errorf("Got %d keys, wanted one", len(keys))
		}
		scw_secret_wrapper.CleanUp(t)

	})

	t.Cleanup(func() {
		scw_secret_wrapper.CleanUp(t)
	})

}
