package scw_secret_manager

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"os"
	"reflect"
	"strconv"
	"testing"

	"github.com/joho/godotenv"
)

//TODO: Consider edge cases

func TestClientCreation(t *testing.T) {
	godotenv.Load(".env")
	t.Run("with 'manual loading' ", func(t *testing.T) {
		var setupData = ScaleWaySetupData{ProjectID: os.Getenv("SCW_DEFAULT_PROJECT_ID"), AccessKey: os.Getenv("SCW_ACCESS_KEY"), SecretKey: os.Getenv("SCW_SECRET_KEY"), Region: os.Getenv("SCW_DEFAULT_REGION")}
		// Would panic if the client creation fails
		NewScaleWayWrapper(setupData)
	})
	t.Run("with 'automatic loading' ", func(t *testing.T) {
		NewScaleWayWrapperFromEnv()
	})
}

func TestSecretCreation(t *testing.T) {

	wrapper := NewScaleWayWrapperFromEnv()
	input := bytes.NewBufferString("test")
	secret, err := wrapper.SetSecret("test", input.Bytes())
	if err != nil {
		t.Error(err)
		t.Logf("THIS IS OK IF YOU DONT CONNECT TO THE SCALEWAY CONSOLE")
	}
	if secret.ID == "" {
		t.Errorf("Secret is nil")
	}

	t.Cleanup(func() {
		CleanUp(t, []string{secret.ID})
	})
}

func TestListSecrets(t *testing.T) {
	t.Run("check if secrets exists", func(t *testing.T) {

		want := 4 // PRIVATE_KEY, apiKey, ENCRYPTION_KEY, test key
		wrapper := NewScaleWayWrapperFromEnv()
		input := bytes.NewBufferString("test")
		secret, err := wrapper.SetSecret("test", input.Bytes())
		if err != nil {
			t.Error(err)
		}
		if secrets, err := wrapper.ListSecrets(); err != nil {
			t.Error(err)
		} else {
			if secrets.TotalCount != uint32(want) {
				t.Errorf("Got %d secrets, wanted %d", secrets.TotalCount, want)
			}

		}
		t.Cleanup(func() {
			CleanUp(t, []string{secret.ID})
		})
	})
}

// Tests if the secret creation fails if the secret name is too short
func TestFailingSecretCreation(t *testing.T) {
	wrapper := NewScaleWayWrapperFromEnv()
	input := bytes.NewBufferString("b")
	_, err := wrapper.SetSecret("a", input.Bytes())
	if err == nil {
		t.Error(err)
	}

	//t.Cleanup(func() { CleanUp(t, []string{secret.ID}) })
}

func TestCreateSecretVersion(t *testing.T) {
	///Tests if the secret version is created for a particular secret
	wrapper := NewScaleWayWrapperFromEnv()
	input := bytes.NewBufferString("c")
	secret, err := wrapper.SetSecret("testSecret", input.Bytes())
	if err != nil {
		t.Error(err)
	}

	if err := wrapper.CreateNewSecretVersion(secret, []byte("c")); err != nil {
		t.Error(err)
	}
	if versionHolder, err := wrapper.ListSecretVersions(secret.ID); err != nil {
		t.Error(err)
	} else {
		if versionHolder.TotalCount != 2 {
			t.Errorf("Got %d secrets, wanted %d", versionHolder.TotalCount, 2)
		}

		t.Cleanup(func() { CleanUp(t, []string{secret.ID}) })
	}
}

func TestListSecretsWName(t *testing.T) {
	godotenv.Load(".env")
	wrapper := NewScaleWayWrapperFromEnv()
	input1 := bytes.NewBufferString("test")
	input2 := bytes.NewBufferString("test2")
	secret1, err := wrapper.SetSecret("tester", input1.Bytes())
	if err != nil {
		t.Error(err)
	}
	secret2, err := wrapper.SetSecret("tester2", input2.Bytes())
	if err != nil {
		t.Error(err)
	}
	holder, err := wrapper.ListSecrets("tester")
	if err != nil {
		t.Error(err)
	}
	want := 1
	if int(holder.TotalCount) != want {
		t.Errorf("Got %d secrets, wanted %d", holder.TotalCount, want)
	}
	t.Cleanup(func() { CleanUp(t, []string{secret1.ID, secret2.ID}) })
}

func TestValues(t *testing.T) {
	t.Run("test with string", func(*testing.T) {
		wrapper := NewScaleWayWrapperFromEnv()
		const want = "test"
		input := bytes.NewBufferString("test")
		secret, err := wrapper.SetSecret("test3", input.Bytes())
		if err != nil {
			t.Error(err)
		}
		secretHolder, err := wrapper.ListSecrets("test3")
		if err != nil {
			t.Error(err)
		} else if secretHolder.TotalCount != 1 {
			t.Error("Got more or less than one secret")
		}
		secretVersions, err := wrapper.ListSecretVersions(secret.ID)
		if err != nil {
			t.Error(err)
		}
		if secretVersions.TotalCount != 1 {
			t.Error("Got more or less than one secret version")
		}
		secretV := secretVersions.SecretVersions[0]
		revision := strconv.FormatUint(uint64(secretV.Revision), 10)
		data, err := wrapper.GetSecretData("test3", revision)
		if err != nil {
			t.Error(err)
		}
		t.Log(string(data))
		if string(data) != want {
			t.Errorf("Got %s, wanted %s", string(data), want)
		}
		t.Cleanup(func() { CleanUp(t, []string{secret.ID}) })
	})
	t.Run("test with zertifikate key", func(*testing.T) {
		godotenv.Load(".env")
		var setupData = ScaleWaySetupData{ProjectID: os.Getenv("SCW_DEFAULT_PROJECT_ID"), AccessKey: os.Getenv("SCW_ACCESS_KEY"), SecretKey: os.Getenv("SCW_SECRET_KEY"), Region: os.Getenv("SCW_DEFAULT_REGION")}
		wrapper := NewScaleWayWrapper(setupData)

		generatedKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
		if err != nil {
			t.Error(err)
		}
		bytes, err := x509.MarshalECPrivateKey(generatedKey)
		if err != nil {
			t.Error(err)
		}
		privPemBlock := &pem.Block{
			Type:  "EC PRIVATE KEY",
			Bytes: bytes,
		}
		encodedBytes := pem.EncodeToMemory(privPemBlock)
		if err != nil {
			t.Error(err)
		}
		secret, err := wrapper.SetSecret("keyTest", encodedBytes)
		if err != nil {
			t.Error(err)
		}
		secretHolder, err := wrapper.ListSecrets("keyTest")
		if err != nil {
			t.Error(err)
		} else if secretHolder.TotalCount != 1 {
			t.Error("Got more or less than one secret")
		}
		secretVersions, err := wrapper.ListSecretVersions(secret.ID)
		if err != nil {
			t.Error(err)
		}
		if secretVersions.TotalCount != 1 {
			t.Error("Got more or less than one secret version")
		}
		secretV := secretVersions.SecretVersions[0]
		revision := strconv.FormatUint(uint64(secretV.Revision), 10)
		data, err := wrapper.GetSecretData("keyTest", revision)
		if err != nil {
			t.Error(err)
		}
		bloc, _ := pem.Decode(data)
		if reflect.DeepEqual(bloc.Bytes, bytes) == false {
			t.Errorf("Got %s, wanted %s", string(data), bytes)
		}
		t.Cleanup(func() { CleanUp(t, []string{secret.ID}) })
	})
}

func TestSecretAccessAfterDeletion(t *testing.T) {
	t.Run("test with string", func(*testing.T) {
		client := NewScaleWayWrapperFromEnv()
		secret, err := client.SetSecret("testKey", []byte("testValue"))
		if err != nil {
			t.Error(err)
		}
		err = client.CreateNewSecretVersion(secret, []byte("testValue2"))
		if err != nil {
			t.Error(err)
		}
		secrets, err := client.ListSecrets("testKey")
		if err != nil {
			t.Error(err)
		}
		secretVersionHolder, err := client.ListSecretVersions(secrets.Secrets[0].ID)
		if (err) != nil {
			t.Error(err)
		}
		versions := secretVersionHolder.SecretVersions
		err = client.DeleteSecretVersion(secrets.Secrets[0].ID, strconv.FormatUint(uint64(versions[1].Revision), 10))
		if err != nil {
			t.Error(err)
		}
		secretVersions, err := client.ListSecretVersions(secrets.Secrets[0].ID)
		if err != nil {
			t.Log(err)
		}
		var apiKeys []string
		for _, secretVersion := range secretVersions.SecretVersions {
			if secretVersion.Status != "destroyed" {
				data, err := client.GetSecretData("testKey", strconv.FormatUint(uint64(secretVersion.Revision), 10))
				apiKeys = append(apiKeys, string(data))
				if err != nil {
					/// Should also be impossible to not get the data if the rest is true
					t.Error(err)
				}
			}
		}
		if len(apiKeys) != 1 {
			t.Errorf("Got %d api keys, wanted %d", len(apiKeys), 1)
		}
		t.Cleanup(func() { CleanUp(t, []string{secret.ID}) })

	})
}
