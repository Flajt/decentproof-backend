package scw_secret_manager

import (
	"encoding/base64"
	"os"
	"time"
)

type LocalScaleWayWrapper struct {
	keyStore map[string][]byte
}

// Loads the following environment variables:
// - ENCRYPTION_KEY base64 encoded encryption key 32 bytes long
// - apiKey base64 encoded api key 32 bytes long
// - PRIVATE_KEY ECDSA Private Key in PEM format
func NewScaleWayWrapperForDev() IScaleWayWrapper {
	keyStore := make(map[string][]byte)
	encryptionKey := os.Getenv("ENCRYPTION_KEY")
	apiKey := os.Getenv("apiKey")
	privateKey := os.Getenv("PRIVATE_KEY")
	if encryptionKey != "" {
		decoded, err := base64.StdEncoding.DecodeString(encryptionKey)
		if err != nil {
			panic(err)
		}
		keyStore["ENCRYPTION_KEY"] = decoded
	}
	if apiKey != "" {
		keyStore["apiKey"] = []byte(apiKey)
	}
	if privateKey != "" {
		keyStore["PRIVATE_KEY"] = []byte(privateKey)
	}

	return &LocalScaleWayWrapper{keyStore: keyStore}

}

func (scalewayWrapper *LocalScaleWayWrapper) ListSecrets(names ...string) (SecretHolder, error) {
	keys := make([]string, 0, len(scalewayWrapper.keyStore))
	secrets := make([]*Secret, 0, len(scalewayWrapper.keyStore))
	for k := range scalewayWrapper.keyStore {
		keys = append(keys, k)
		now := time.Now()
		secrets = append(secrets, &Secret{ID: k, Name: k, CreatedAt: &now})
	}

	return SecretHolder{TotalCount: uint32(len(keys)), Secrets: secrets}, nil
}

func (ScalewayWrapper *LocalScaleWayWrapper) ListSecretVersions(secretID string) (SecretVersionHolder, error) {
	if _, ok := ScalewayWrapper.keyStore[secretID]; ok {
		now := time.Now()
		return SecretVersionHolder{TotalCount: 1, SecretVersions: []SecretVersion{{CreatedAt: &now, SecretID: secretID, IsLatest: true}}}, nil
	} else {
		return SecretVersionHolder{}, nil
	}
}

func (scalewayWrapper *LocalScaleWayWrapper) GetSecretData(secretName string, revision string) ([]byte, error) {
	if secret, ok := scalewayWrapper.keyStore[secretName]; ok {
		return secret, nil
	} else {
		return []byte{}, nil
	}
}

func (scalewayWrapper *LocalScaleWayWrapper) SetSecret(secretName string, secretValue []byte) (Secret, error) {
	scalewayWrapper.keyStore[secretName] = secretValue
	now := time.Now()
	return Secret{ID: secretName, Name: secretName, CreatedAt: &now}, nil

}

func (scalewayWrapper *LocalScaleWayWrapper) CreateNewSecretVersion(secret Secret, data []byte) error {
	return nil
}

func (scalewayWrapper *LocalScaleWayWrapper) DeleteSecret(id string) error {
	//Delete entry from map
	delete(scalewayWrapper.keyStore, id)
	return nil

}

func (scalewayWrapper *LocalScaleWayWrapper) DeleteSecretVersion(id string, revision string) error {
	return nil
}
