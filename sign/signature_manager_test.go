package sign

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"testing"

	scw_secret_wrapper "github.com/Flajt/decentproof-backend/scw_secret_wrapper"
	"go.uber.org/mock/gomock"
)

func TestInitalisation(t *testing.T) {

	t.Run("working", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		privKey, err := generatePrivKey(t)
		if err != nil {
			t.Errorf("Unexpected error, got %v", err)
		}
		wrapper := scw_secret_wrapper.NewMockIScaleWayWrapper(ctrl)
		wrapper.EXPECT().ListSecrets("PRIVATE_KEY").Return(scw_secret_wrapper.SecretHolder{Secrets: []*scw_secret_wrapper.Secret{{ID: "test", Name: "test"}}}, nil)
		wrapper.EXPECT().ListSecretVersions(gomock.Any()).Return(scw_secret_wrapper.SecretVersionHolder{SecretVersions: []scw_secret_wrapper.SecretVersion{{Revision: 1, IsLatest: true}}}, nil)
		wrapper.EXPECT().GetSecretData(gomock.Any(), gomock.Any()).Return(privKey, nil)
		manager := NewSignatureManager(wrapper)
		err = manager.InitSignatureManager()
		if err != nil {
			t.Error(err)
		}
		//defer scw_secret_wrapper.CleanUp(t)

	})
	t.Run("signing hash", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		privKey, err := generatePrivKey(t)
		if err != nil {
			t.Errorf("Unexpected error, got %v", err)
		}
		wrapper := scw_secret_wrapper.NewMockIScaleWayWrapper(ctrl)
		wrapper.EXPECT().ListSecrets("PRIVATE_KEY").Return(scw_secret_wrapper.SecretHolder{Secrets: []*scw_secret_wrapper.Secret{{ID: "test", Name: "test"}}}, nil)
		wrapper.EXPECT().ListSecretVersions(gomock.Any()).Return(scw_secret_wrapper.SecretVersionHolder{SecretVersions: []scw_secret_wrapper.SecretVersion{{Revision: 1, IsLatest: true}}}, nil)
		wrapper.EXPECT().GetSecretData(gomock.Any(), gomock.Any()).Return(privKey, nil)
		signatureManager := NewSignatureManager(wrapper)
		err = signatureManager.InitSignatureManager()
		if err != nil {
			t.Fatal(err)
		}
		signature, err := signatureManager.SignData([]byte("test"))
		if err != nil {
			t.Error(err)
		}
		if signature == nil {
			t.Errorf("Expected signature, got nil")
		}
		//sha256Hash := sha256.Sum256([]byte("test"))
		//sha256HashSlice := sha256Hash[:]
		isValid := signatureManager.VerifyData([]byte("test"), signature)
		if isValid != true {
			t.Errorf("Expected true, got %v", isValid)
		}
		//defer scw_secret_wrapper.CleanUp(t)

	})
	//t.Cleanup(func() { scw_secret_wrapper.CleanUp(t) })

}

func generatePrivKey(t *testing.T) ([]byte, error) {
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
	bytes, err := x509.MarshalECPrivateKey(key)
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
	bloc := &pem.Block{Type: "EC PRIVATE KEY", Bytes: bytes}
	privKey := pem.EncodeToMemory(bloc)
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
	return privKey, nil
}
