package decentproof_cronjob

import (
	"testing"

	decentproof_cronjob "github.com/decentproof-cron"
)

//TODO: Consider edge cases

func TestClientCreation(t *testing.T) {
	if _, err := decentproof_cronjob.NewScaleWayWrapper(); err != nil {
		t.Error(err)
	}
}

func TestSecretCreation(t *testing.T) {
	if wrapper, err := decentproof_cronjob.NewScaleWayWrapper(); err != nil {
		t.Error(err)
	} else {
		if err := wrapper.SetSecret("test", "test"); err != nil {
			t.Error(err)
			t.Logf("THIS IS OK IF YOU DONT CONNECT TO THE SCALEWAY CONSOLE")
		}
	}
	t.Cleanup(func() {
		cleanUp(t)
	})
}

func TestListSecrets(t *testing.T) {
	want := 1
	if wrapper, err := decentproof_cronjob.NewScaleWayWrapper(); err != nil {
		t.Error(err)
	} else {
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
	}
	t.Cleanup(func() { cleanUp(t) })

}

// Tests if the secret creation fails if the secret name is too short
func TestFailingSecretCreation(t *testing.T) {
	if wrapper, err := decentproof_cronjob.NewScaleWayWrapper(); err != nil {
		t.Error(err)
	} else {
		if err := wrapper.SetSecret("a", "b"); err == nil {
			t.Error(err)
		}
	}
	t.Cleanup(func() { cleanUp(t) })
}

func TestCreateSecretVersion(t *testing.T) {
	///Tests if the secret version is created for a particular secret
	if wrapper, err := decentproof_cronjob.NewScaleWayWrapper(); err != nil {
		t.Error(err)
	} else {
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
		}
		t.Cleanup(func() { cleanUp(t) })
	}
}

func cleanUp(t *testing.T) {
	if wrapper, err := decentproof_cronjob.NewScaleWayWrapper(); err != nil {
		t.Error(err)
		t.Log("MANUAL CLEAN UP REQUIRED !!!!!!")
	} else {
		if secrets, err := wrapper.ListSecrets(); err != nil {
			t.Error(err)
			t.Log("MANUAL CLEAN UP REQUIRED !!!!!!")
		} else {
			for _, secret := range secrets.Secrets {
				if err := wrapper.DeleteSecret(secret.ID); err != nil {
					t.Error(err)
					t.Log("MANUAL CLEAN UP REQUIRED !!!!!!")
				}
			}
		}

	}
}
