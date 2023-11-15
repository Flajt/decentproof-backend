package encryption_service

import (
	"encoding/base64"
	"testing"
	"time"

	"github.com/Flajt/decentproof-backend/helper"
	scw_secret_manager "github.com/Flajt/decentproof-backend/scw_secret_wrapper"
	"go.uber.org/mock/gomock"
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

	t.Run("correct key should decode data correctly", func(*testing.T) {
		ctrl := gomock.NewController(t)
		m := scw_secret_manager.NewMockIScaleWayWrapper(ctrl)
		defer ctrl.Finish()
		base64Key := helper.GenerateApiKey(32)
		bytes, err := base64.StdEncoding.DecodeString(base64Key)
		if err != nil {
			t.Error(err)
		}
		//scwWrapper.SetSecret("ENCRYPTION_KEY", bytes)
		m.EXPECT().ListSecrets(gomock.Any()).Return(scw_secret_manager.SecretHolder{Secrets: []*scw_secret_manager.Secret{{ID: "test", Name: "test"}}, TotalCount: 1}, nil).MaxTimes(2)
		m.EXPECT().ListSecretVersions(gomock.Any()).Return(scw_secret_manager.SecretVersionHolder{SecretVersions: []scw_secret_manager.SecretVersion{{Revision: 1, CreatedAt: &time.Time{}}}, TotalCount: 1}, nil).MaxTimes(2)
		m.EXPECT().GetSecretData(gomock.Any(), gomock.Any()).Return(bytes, nil).MaxTimes(2)
		service := NewEncryptionService(m)
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
		ctrl := gomock.NewController(t)
		m := scw_secret_manager.NewMockIScaleWayWrapper(ctrl)
		defer ctrl.Finish()

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
		//scwWrapper.SetSecret("ENCRYPTION_KEY", bytes1)
		service := NewEncryptionServiceFromKey(bytes1)
		encryptedData, err := service.EncryptData([]byte("test"))
		if err != nil {
			t.Error(err)
		}
		m.EXPECT().ListSecrets("ENCRYPTION_KEY").Return(scw_secret_manager.SecretHolder{Secrets: []*scw_secret_manager.Secret{{ID: "test", Name: "test"}}, TotalCount: 1}, nil)
		m.EXPECT().ListSecretVersions(gomock.Any()).Return(scw_secret_manager.SecretVersionHolder{SecretVersions: []scw_secret_manager.SecretVersion{{Revision: 1, CreatedAt: &time.Time{}}}, TotalCount: 1}, nil)
		m.EXPECT().GetSecretData("test", "1").Return(bytes2, nil)
		decryptionService := NewEncryptionService(m)
		decryptionService.DecryptData(encryptedData.Data, encryptedData.Nonce)
	})

}

func TestMultipleKeyHandleing(t *testing.T) {

	t.Run("should use second key if the new one doesn't work", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		m := scw_secret_manager.NewMockIScaleWayWrapper(ctrl)

		defer ctrl.Finish()

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

		now := time.Now()
		later := now.Add(1 * time.Second)
		m.EXPECT().ListSecrets(gomock.Any()).Return(scw_secret_manager.SecretHolder{Secrets: []*scw_secret_manager.Secret{{ID: "test", Name: "test"}}, TotalCount: 1}, nil)
		m.EXPECT().ListSecretVersions(gomock.Any()).Return(scw_secret_manager.SecretVersionHolder{SecretVersions: []scw_secret_manager.SecretVersion{{Revision: 1, CreatedAt: &now}, {Revision: 2, CreatedAt: &later}}, TotalCount: 2}, nil)
		m.EXPECT().GetSecretData(gomock.Any(), "1").Return(firstKeyBytes, nil)
		m.EXPECT().GetSecretData("test", "2").Return(secondKeyBytes, nil)

		service := NewEncryptionServiceFromKey(firstKeyBytes)
		encryptedData, err := service.EncryptData([]byte("test"))
		if err != nil {
			t.Error(err)
		}
		service = NewEncryptionService(m)

		//secret, err := scwWrapper.SetSecret("ENCRYPTION_KEY", firstKeyBytes)

		time.Sleep(10 * time.Millisecond)
		//err = scwWrapper.CreateNewSecretVersion(secret, secondKeyBytes)

		decryptedData, err := service.DecryptData(encryptedData.Data, encryptedData.Nonce)
		if err != nil {
			t.Error(err)
		}
		if string(decryptedData) != "test" {
			t.Error("Decrypted data does not match original data")
		}

	})
}
