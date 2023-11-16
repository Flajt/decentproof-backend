package sign

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"testing"

	scw_secret_wrapper "github.com/Flajt/decentproof-backend/scw_secret_wrapper"
)

func TestInitalisation(t *testing.T) {

	t.Run("working", func(t *testing.T) {
		privKey, err := generatePrivKey(t)
		if err != nil {
			t.Errorf("Unexpected error, got %v", err)
		}
		wrapper := scw_secret_wrapper.NewScaleWayWrapperForDev(true)
		_, err = wrapper.SetSecret("PRIVATE_KEY", privKey)
		if err != nil {
			t.Errorf("Unexpected error, got %v", err)
		}
		manager := NewSignatureManager(wrapper)
		err = manager.InitSignatureManager()
		if err != nil {
			t.Error(err)
		}
		//defer scw_secret_wrapper.CleanUp(t)

	})
	t.Run("signing hash", func(t *testing.T) {
		privKey, err := generatePrivKey(t)
		if err != nil {
			t.Errorf("Unexpected error, got %v", err)
		}
		wrapper := scw_secret_wrapper.NewScaleWayWrapperForDev(true)
		secret, err := wrapper.SetSecret("PRIVATE_KEY", privKey)
		t.Log(secret.ID)

		if err != nil {
			t.Errorf("Unexpected error, got %v", err)
		}
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
