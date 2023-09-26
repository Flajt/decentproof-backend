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
	godotenv.Load(".env")
	var setupData = ScaleWaySetupData{ProjectID: os.Getenv("SCW_DEFAULT_PROJECT_ID"), AccessKey: os.Getenv("SCW_ACCESS_KEY"), SecretKey: os.Getenv("SCW_SECRET_KEY"), Region: os.Getenv("SCW_DEFAULT_REGION")}

	wrapper := NewScaleWayWrapper(setupData)
	input := bytes.NewBufferString("test")
	secret, err := wrapper.SetSecret("test", input.Bytes())
	if err != nil {
		t.Error(err)
		t.Logf("THIS IS OK IF YOU DONT CONNECT TO THE SCALEWAY CONSOLE")
	}
	if secret == nil {
		t.Errorf("Secret is nil")
	}

	t.Cleanup(func() {
		CleanUp(t)
	})
}

func TestListSecrets(t *testing.T) {
	godotenv.Load(".env")
	var setupData = ScaleWaySetupData{ProjectID: os.Getenv("SCW_DEFAULT_PROJECT_ID"), AccessKey: os.Getenv("SCW_ACCESS_KEY"), SecretKey: os.Getenv("SCW_SECRET_KEY"), Region: os.Getenv("SCW_DEFAULT_REGION")}

	want := 1
	wrapper := NewScaleWayWrapper(setupData)
	input := bytes.NewBufferString("test")
	if _, err := wrapper.SetSecret("test", input.Bytes()); err != nil {
		t.Error(err)
	}
	if secrets, err := wrapper.ListSecrets(); err != nil {
		t.Error(err)
	} else {
		if secrets.TotalCount != uint32(want) {
			t.Errorf("Got %d secrets, wanted %d", secrets.TotalCount, want)
		}

	}
	t.Cleanup(func() { CleanUp(t) })

}

// Tests if the secret creation fails if the secret name is too short
func TestFailingSecretCreation(t *testing.T) {
	godotenv.Load(".env")
	var setupData = ScaleWaySetupData{ProjectID: os.Getenv("SCW_DEFAULT_PROJECT_ID"), AccessKey: os.Getenv("SCW_ACCESS_KEY"), SecretKey: os.Getenv("SCW_SECRET_KEY"), Region: os.Getenv("SCW_DEFAULT_REGION")}
	wrapper := NewScaleWayWrapper(setupData)
	input := bytes.NewBufferString("b")
	if _, err := wrapper.SetSecret("a", input.Bytes()); err == nil {
		t.Error(err)
	}

	t.Cleanup(func() { CleanUp(t) })
}

func TestCreateSecretVersion(t *testing.T) {
	godotenv.Load(".env")
	///Tests if the secret version is created for a particular secret
	var setupData = ScaleWaySetupData{ProjectID: os.Getenv("SCW_DEFAULT_PROJECT_ID"), AccessKey: os.Getenv("SCW_ACCESS_KEY"), SecretKey: os.Getenv("SCW_SECRET_KEY"), Region: os.Getenv("SCW_DEFAULT_REGION")}

	wrapper := NewScaleWayWrapper(setupData)
	input := bytes.NewBufferString("c")
	secret, err := wrapper.SetSecret("testSecret", input.Bytes())
	if err != nil {
		t.Error(err)
	}

	if err := wrapper.CreateNewSecretVersion(*secret, []byte("c")); err != nil {
		t.Error(err)
	}
	if versionHolder, err := wrapper.ListSecretVersions(secret.ID); err != nil {
		t.Error(err)
	} else {
		if versionHolder.TotalCount != 2 {
			t.Errorf("Got %d secrets, wanted %d", versionHolder.TotalCount, 2)
		}

		t.Cleanup(func() { CleanUp(t) })
	}
}

func TestListSecretsWName(t *testing.T) {
	godotenv.Load(".env")
	var setupData = ScaleWaySetupData{ProjectID: os.Getenv("SCW_DEFAULT_PROJECT_ID"), AccessKey: os.Getenv("SCW_ACCESS_KEY"), SecretKey: os.Getenv("SCW_SECRET_KEY"), Region: os.Getenv("SCW_DEFAULT_REGION")}
	wrapper := NewScaleWayWrapper(setupData)
	input1 := bytes.NewBufferString("test")
	input2 := bytes.NewBufferString("test2")
	if _, err := wrapper.SetSecret("tester", input1.Bytes()); err != nil {
		t.Error(err)
	}
	if _, err := wrapper.SetSecret("tester2", input2.Bytes()); err != nil {
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
	t.Cleanup(func() { CleanUp(t) })
}

func TestValues(t *testing.T) {
	t.Run("test with string", func(*testing.T) {
		godotenv.Load(".env")
		var setupData = ScaleWaySetupData{ProjectID: os.Getenv("SCW_DEFAULT_PROJECT_ID"), AccessKey: os.Getenv("SCW_ACCESS_KEY"), SecretKey: os.Getenv("SCW_SECRET_KEY"), Region: os.Getenv("SCW_DEFAULT_REGION")}
		wrapper := NewScaleWayWrapper(setupData)
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
	})
	t.Cleanup(func() { CleanUp(t) })
}
