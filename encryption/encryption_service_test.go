package encryption_service

import (
	"encoding/base64"
	"testing"
	"time"

	"github.com/Flajt/decentproof-backend/helper"
	scw_secret_manager "github.com/Flajt/decentproof-backend/scw_secret_wrapper"
)

func TestEncryptionWLocalKey(t *testing.T) {
	t.Run("correct key should decode data correctly", func(*testing.T) {
		base64Key := helper.GenerateApiKey(32)
		bytes, err := base64.StdEncoding.DecodeString(base64Key)
		if err != nil {
			t.Error(err)
		}
		service := NewEncryptionServiceFromKey(bytes)
		encryptedData, err := service.EncryptData([]byte("test"))
		if err != nil {
			t.Error(err)
		}
		decryptedData, err := service.DecryptData(encryptedData.Data, encryptedData.Nonce)
		if err != nil {
			t.Error(err)
		}
		if string(decryptedData) != "test" {
			t.Error("Decrypted data does not match original data")
		}
	})
	t.Run("incorrect key should return error", func(*testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()
		base64Key1 := helper.GenerateApiKey(32)
		base64Key2 := helper.GenerateApiKey(32)
		bytes1, err := base64.StdEncoding.DecodeString(base64Key1)
		if err != nil {
			t.Error(err)
		}
		bytes2, err := base64.StdEncoding.DecodeString(base64Key2)
		if err != nil {
			t.Error(err)
		}
		service := NewEncryptionServiceFromKey(bytes1)
		encryptedData, err := service.EncryptData([]byte("test"))
		if err != nil {
			t.Error(err)
		}
		service2 := NewEncryptionServiceFromKey(bytes2)
		service2.DecryptData(encryptedData.Data, encryptedData.Nonce)

	})

}

func TestWithSCWStoredKey(t *testing.T) {
	scwWrapper := scw_secret_manager.NewScaleWayWrapperFromEnv()

	t.Run("correct key should decode data correctly", func(*testing.T) {
		base64Key := helper.GenerateApiKey(32)
		bytes, err := base64.StdEncoding.DecodeString(base64Key)
		if err != nil {
			t.Error(err)
		}
		scwWrapper.SetSecret("ENCRYPTION_KEY", bytes)
		service := NewEncryptionService(*scwWrapper)
		encryptedData, err := service.EncryptData([]byte("test"))
		if err != nil {
			t.Error(err)
		}
		decryptedData, err := service.DecryptData(encryptedData.Data, encryptedData.Nonce)
		if err != nil {
			t.Error(err)
		}
		if string(decryptedData) != "test" {
			t.Error("Decrypted data does not match original data")
		}
	})

	t.Run("incorrect key should return error", func(*testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()
		base64Key1 := helper.GenerateApiKey(32)
		base64Key2 := helper.GenerateApiKey(32)
		bytes1, err := base64.StdEncoding.DecodeString(base64Key1)
		if err != nil {
			t.Error(err)
		}
		bytes2, err := base64.StdEncoding.DecodeString(base64Key2)
		if err != nil {
			t.Error(err)
		}
		scwWrapper.SetSecret("ENCRYPTION_KEY", bytes1)
		service := NewEncryptionService(*scwWrapper)
		encryptedData, err := service.EncryptData([]byte("test"))
		if err != nil {
			t.Error(err)
		}
		cleanUp(scwWrapper, t)
		scwWrapper.SetSecret("ENCRYPTION_KEY", bytes2)
		service.DecryptData(encryptedData.Data, encryptedData.Nonce)
	})
	t.Cleanup(func() {
		cleanUp(scwWrapper, t)
	})

}
func cleanUp(scwWrapper *scw_secret_manager.ScalewayWrapper, t *testing.T) {
	secrets, err := scwWrapper.ListSecrets("ENCRYPTION_KEY")
	if err != nil {
		t.Error(err)
	}
	for _, secret := range secrets.Secrets {
		scwWrapper.DeleteSecret(secret.ID)
	}
}

func TestMultipleKeyHandleing(t *testing.T) {
	scwWrapper := scw_secret_manager.NewScaleWayWrapperFromEnv()

	firstKey := helper.GenerateApiKey(32)
	secondKey := helper.GenerateApiKey(32)

	firstKeyBytes, err := base64.StdEncoding.DecodeString(firstKey)
	if err != nil {
		t.Error(err)
	}
	secondKeyBytes, err := base64.StdEncoding.DecodeString(secondKey)
	if err != nil {
		t.Error(err)
	}

	_, err = scwWrapper.SetSecret("ENCRYPTION_KEY", firstKeyBytes)
	if err != nil {
		t.Error(err)
	}
	service := NewEncryptionService(*scwWrapper)
	encryptedData, err := service.EncryptData([]byte("test"))
	if err != nil {
		t.Error(err)
	}
	cleanUp(scwWrapper, t)

	secret, err := scwWrapper.SetSecret("ENCRYPTION_KEY", secondKeyBytes)
	if err != nil {
		t.Error(err)
	}
	time.Sleep(10 * time.Millisecond)
	err = scwWrapper.CreateNewSecretVersion(*secret, firstKeyBytes)
	if err != nil {
		t.Error(err)
	}
	decryptedData, err := service.DecryptData(encryptedData.Data, encryptedData.Nonce)
	if err != nil {
		t.Error(err)
	}
	if string(decryptedData) != "test" {
		t.Error("Decrypted data does not match original data")
	}

	t.Cleanup(func() {
		cleanUp(scwWrapper, t)
	})
}
