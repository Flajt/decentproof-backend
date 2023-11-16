package scw_secret_manager

import (
	"strconv"
	"testing"
)

func TestScwWrapperForDev(t *testing.T) {

	t.Run("should be empty if initEmpty it passed secrets", func(t *testing.T) {
		wrapper := NewScaleWayWrapperForDev(true)
		secrets, _ := wrapper.ListSecrets()
		if len(secrets.Secrets) != 0 {
			t.Errorf("Expected 0 secrets, got %d", len(secrets.Secrets))
		}
	})
	t.Run("should return the correct secret", func(t *testing.T) {
		wrapper := NewScaleWayWrapperForDev(true)
		wrapper.SetSecret("ENCRYPTION_KEY", []byte("test"))
		secrets, _ := wrapper.ListSecrets("ENCRYPTION_KEY")
		if len(secrets.Secrets) != 1 {
			t.Errorf("Expected 1 secret, got %d", len(secrets.Secrets))
		}
		if secrets.Secrets[0].Name != "ENCRYPTION_KEY" {
			t.Errorf("Expected ENCRYPTION_KEY, got %s", secrets.Secrets[0].Name)
		}
	})
	t.Run("should list all secrets if no search term is added", func(t *testing.T) {
		wrapper := NewScaleWayWrapperForDev(true)
		wrapper.SetSecret("ENCRYPTION_KEY", []byte("test"))
		wrapper.SetSecret("SUPERDUPER", []byte("test"))
		secrets, _ := wrapper.ListSecrets()
		if len(secrets.Secrets) != 2 {
			t.Errorf("Expected 2 secrets, got %d", len(secrets.Secrets))
		}
	})
	t.Run("should return the correct secret version", func(t *testing.T) {
		wrapper := NewScaleWayWrapperForDev(true)
		wrapper.SetSecret("ENCRYPTION_KEY", []byte("test"))
		secret, _ := wrapper.ListSecrets("ENCRYPTION_KEY")
		secretVersion, _ := wrapper.ListSecretVersions(secret.Secrets[0].ID)
		if len(secretVersion.SecretVersions) != 1 {
			t.Errorf("Expected 1 secret version, got %d", len(secretVersion.SecretVersions))
		}
		if secretVersion.SecretVersions[0].IsLatest != true {
			t.Errorf("Expected true, got %v", secretVersion.SecretVersions[0].IsLatest)
		}
		if secretVersion.SecretVersions[0].Revision != 1 {
			t.Errorf("Expected 1, got %v", secretVersion.SecretVersions[0].Revision)
		}
	})
	t.Run("should return the correct secret data", func(t *testing.T) {
		wrapper := NewScaleWayWrapperForDev(true)
		wrapper.SetSecret("ENCRYPTION_KEY", []byte("test"))
		secrets, _ := wrapper.ListSecrets("ENCRYPTION_KEY")
		secretVersion, _ := wrapper.ListSecretVersions(secrets.Secrets[0].ID)
		secretData, _ := wrapper.GetSecretData(secrets.Secrets[0].Name, strconv.FormatUint(uint64(secretVersion.SecretVersions[0].Revision), 10))
		if string(secretData) != "test" {
			t.Errorf("Expected test, got %s", string(secretData))
		}
	})
	t.Run("should delete data", func(t *testing.T) {
		wrapper := NewScaleWayWrapperForDev(true)
		secret, _ := wrapper.SetSecret("ENCRYPTION_KEY", []byte("test"))
		wrapper.DeleteSecret(secret.ID)
		secrets, _ := wrapper.ListSecrets()
		if len(secrets.Secrets) != 0 {
			t.Errorf("Expected 0 secrets, got %d", len(secrets.Secrets))
		}
	})
}
