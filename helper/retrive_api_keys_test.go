package helper

import (
	"testing"
	"time"

	scw_secret_manager "github.com/Flajt/decentproof-backend/scw_secret_wrapper"
	"github.com/joho/godotenv"
	"go.uber.org/mock/gomock"
)

func TestRetrievApiKeys(t *testing.T) {
	godotenv.Load("../.env")
	t.Run("with zero entries", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		m := scw_secret_manager.NewMockIScaleWayWrapper(ctrl)
		m.EXPECT().ListSecrets(gomock.Any()).Return(scw_secret_manager.SecretHolder{Secrets: []*scw_secret_manager.Secret{}, TotalCount: 0}, nil)

		defer func() {
			if r := recover(); r != nil {
				t.Log("Recovered from panic")

			}
		}()
		keys := RetrievApiKeys(m)
		if len(keys) != 0 {
			t.Error("Got keys, wanted none")
		}
	})

	t.Run("with one secret entry", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		m := scw_secret_manager.NewMockIScaleWayWrapper(ctrl)
		m.EXPECT().ListSecrets(gomock.Any()).Return(scw_secret_manager.SecretHolder{Secrets: []*scw_secret_manager.Secret{{ID: "apiKey", Name: "apiKey"}}, TotalCount: 1}, nil)
		m.EXPECT().ListSecretVersions(gomock.Any()).Return(scw_secret_manager.SecretVersionHolder{SecretVersions: []scw_secret_manager.SecretVersion{{Revision: 1, CreatedAt: &time.Time{}}}, TotalCount: 1}, nil).MaxTimes(2)
		m.EXPECT().GetSecretData(gomock.Any(), gomock.Any()).Return([]byte("test"), nil)
		keys := RetrievApiKeys(m)

		if len(keys) != 1 {
			t.Errorf("Got %d keys, wanted one", len(keys))
		}
	})
	//Depreciated
	/*t.Run("with two secret entries", func(t *testing.T) {
		keys := RetrievApiKeys()
		if len(keys) != 2 {
			t.Errorf("Got %d keys, wanted one", len(keys))
		}
		scw_secret_wrapper.CleanUp(t)

	})*/ //Depreciated

}
