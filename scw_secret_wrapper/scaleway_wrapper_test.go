package scw_secret_manager

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

//TODO: Consider edge cases

func TestClientCreation(t *testing.T) {
	godotenv.Load(".env")
	var setupData = ScaleWaySetupData{ProjectID: os.Getenv("SCW_DEFAULT_PROJECT_ID"), AccessKey: os.Getenv("SCW_ACCESS_KEY"), SecretKey: os.Getenv("SCW_SECRET_KEY"), Region: os.Getenv("SCW_DEFAULT_REGION")}
	// Would panic if the client creation fails
	NewScaleWayWrapper(setupData)
}

func TestSecretCreation(t *testing.T) {
	godotenv.Load(".env")
	var setupData = ScaleWaySetupData{ProjectID: os.Getenv("SCW_DEFAULT_PROJECT_ID"), AccessKey: os.Getenv("SCW_ACCESS_KEY"), SecretKey: os.Getenv("SCW_SECRET_KEY"), Region: os.Getenv("SCW_DEFAULT_REGION")}

	wrapper := NewScaleWayWrapper(setupData)
	if err := wrapper.SetSecret("test", "test"); err != nil {
		t.Error(err)
		t.Logf("THIS IS OK IF YOU DONT CONNECT TO THE SCALEWAY CONSOLE")
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
	if err := wrapper.SetSecret("test", "b"); err != nil {
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
	if err := wrapper.SetSecret("a", "b"); err == nil {
		t.Error(err)
	}

	t.Cleanup(func() { CleanUp(t) })
}

func TestCreateSecretVersion(t *testing.T) {
	godotenv.Load(".env")
	///Tests if the secret version is created for a particular secret
	var setupData = ScaleWaySetupData{ProjectID: os.Getenv("SCW_DEFAULT_PROJECT_ID"), AccessKey: os.Getenv("SCW_ACCESS_KEY"), SecretKey: os.Getenv("SCW_SECRET_KEY"), Region: os.Getenv("SCW_DEFAULT_REGION")}

	wrapper := NewScaleWayWrapper(setupData)
	if err := wrapper.SetSecret("testSecret", "b"); err != nil {
		t.Error(err)
	}
	secretHolder, err := wrapper.ListSecrets()
	if err != nil {
		t.Error(err)
	}
	if err := wrapper.CreateNewSecretVersion(*secretHolder.Secrets[0], "c"); err != nil {
		t.Error(err)
	}
	if versionHolder, err := wrapper.ListSecretVersions(secretHolder.Secrets[0].ID); err != nil {
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

	if err := wrapper.SetSecret("tester", "test"); err != nil {
		t.Error(err)
	}
	if err := wrapper.SetSecret("tester2", "test2"); err != nil {
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
