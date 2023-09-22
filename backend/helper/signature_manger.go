package decentproof_functions

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"

	secret_wrapper "github.com/Flajt/decentproof-backend/scw_secret_wrapper"
)

type SignatureManager struct {
	privKey       ecdsa.PrivateKey
	secretManager *secret_wrapper.ScalewayWrapper
}

func (sm *SignatureManager) InitSignatureManager() error {
	secretHolder, err := sm.secretManager.ListSecrets("PRIVATE_KEY")
	if err != nil {
		return err
	}
	encryptionKeySecret := secretHolder.Secrets[0]
	secretVersionHolder, err := sm.secretManager.ListSecretVersions(encryptionKeySecret.ID)
	if err != nil {
		return err
	}
	var encryptionKeyIndex = -2
	for index, secretVersion := range secretVersionHolder.SecretVersions {
		if secretVersion.IsLatest {
			encryptionKeyIndex = index
			break
		}
	}
	encryptionKeyVersionData := secretVersionHolder.SecretVersions[encryptionKeyIndex]
	key, err := sm.secretManager.GetSecretData(encryptionKeySecret.ID, encryptionKeyVersionData.SecretID)
	if err != nil {
		return err
	}
	bloc, _ := pem.Decode(key)
	privKey, err := x509.ParseECPrivateKey(bloc.Bytes)
	if err != nil {
		return err
	}
	sm.privKey = *privKey
	return nil
}
func NewSignatureManager(secretManager *secret_wrapper.ScalewayWrapper) *SignatureManager {
	return &SignatureManager{secretManager: secretManager}
}

func (sm *SignatureManager) SignData(data []byte) ([]byte, error) {
	return ecdsa.SignASN1(rand.Reader, &sm.privKey, data)
}
