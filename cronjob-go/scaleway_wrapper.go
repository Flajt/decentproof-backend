package decentproof_cronjob

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
	secret_manager "github.com/scaleway/scaleway-sdk-go/api/secret/v1alpha1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

type SecretHolder struct {
	Secrets    []*secret_manager.Secret
	TotalCount uint32
}

type SecretVersionHolder struct {
	SecretVersions []*secret_manager.SecretVersion
	TotalCount     uint32
}

type ScalewayWrapper struct {
	Client     scw.Client
	Api        *secret_manager.API
	PROJECT_ID string
}

func NewScaleWayWrapper() (*ScalewayWrapper, error) {
	var envLoadingLocation = ".env"
	currentPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	if strings.Contains(currentPath, "test") {
		envLoadingLocation = "../.env"
	}

	if err := godotenv.Load(envLoadingLocation); err != nil {
		fmt.Println(err)
	}
	accessKey := os.Getenv("SCW_ACCESS_KEY")
	secretKey := os.Getenv("SCW_SECRET_KEY")
	projectID := os.Getenv("SCW_DEFAULT_PROJECT_ID")

	if client, err := scw.NewClient(
		scw.WithAuth(accessKey, secretKey),
		scw.WithDefaultRegion(scw.RegionNlAms),
		scw.WithDefaultProjectID(projectID),
	); err != nil {
		return nil, err
	} else {
		api := secret_manager.NewAPI(client)
		return &ScalewayWrapper{Client: *client, Api: api, PROJECT_ID: projectID}, nil
	}
}

func (scalewayWrapper *ScalewayWrapper) ListSecrets(names ...string) (SecretHolder, error) {
	if len(names) > 0 {
		if secrets, err := scalewayWrapper.Api.ListSecrets(&secret_manager.ListSecretsRequest{ProjectID: &scalewayWrapper.PROJECT_ID, Name: &names[0]}); err != nil {
			return SecretHolder{}, err
		} else {
			return SecretHolder{Secrets: secrets.Secrets, TotalCount: secrets.TotalCount}, nil
		}
	} else {
		if secrets, err := scalewayWrapper.Api.ListSecrets(&secret_manager.ListSecretsRequest{ProjectID: &scalewayWrapper.PROJECT_ID}); err != nil {
			return SecretHolder{}, err
		} else {
			return SecretHolder{Secrets: secrets.Secrets, TotalCount: secrets.TotalCount}, nil
		}
	}

}

func (ScalewayWrapper *ScalewayWrapper) ListSecretVersions(secretID string) (SecretVersionHolder, error) {
	if secrets, err := ScalewayWrapper.Api.ListSecretVersions(&secret_manager.ListSecretVersionsRequest{SecretID: secretID}); err != nil {
		return SecretVersionHolder{}, err
	} else {
		return SecretVersionHolder{SecretVersions: secrets.Versions, TotalCount: secrets.TotalCount}, nil
	}
}

func (scalewayWrapper *ScalewayWrapper) GetSecretData(secretName string, revision string) (string, error) {
	if secret, err := scalewayWrapper.Api.GetSecretVersionByName(&secret_manager.GetSecretVersionByNameRequest{Region: scw.RegionNlAms, Revision: revision, SecretName: secretName}); err != nil {
		return "", err
	} else {
		if secretVersion, err := scalewayWrapper.Api.AccessSecretVersion(&secret_manager.AccessSecretVersionRequest{Region: scw.RegionNlAms, SecretID: secret.SecretID}); err != nil {
			return "", err
		} else {
			return string(secretVersion.Data), nil
		}
	}
}

func (scalewayWrapper *ScalewayWrapper) SetSecret(secretName string, secretValue string) error {
	inputBytes := []byte(secretValue)

	if secret, err := scalewayWrapper.Api.CreateSecret(&secret_manager.CreateSecretRequest{Region: scw.RegionNlAms, Name: secretName, Type: secret_manager.SecretTypeUnknownSecretType}); err != nil {
		return err
	} else {
		if _, err := scalewayWrapper.Api.CreateSecretVersion(&secret_manager.CreateSecretVersionRequest{SecretID: secret.ID, Region: scw.RegionNlAms, Data: inputBytes}); err != nil {
			return err
		}
		return nil
	}
}

func (scalewayWrapper *ScalewayWrapper) CreateNewSecretVersion(secret secret_manager.Secret, data string) error {
	inputBytes := []byte(data)
	if _, err := scalewayWrapper.Api.CreateSecretVersion(&secret_manager.CreateSecretVersionRequest{SecretID: secret.ID, Region: scw.RegionNlAms, Data: inputBytes}); err != nil {
		return err
	}
	return nil
}

func (scalewayWrapper *ScalewayWrapper) DeleteSecret(id string) error {
	if err := scalewayWrapper.Api.DeleteSecret(&secret_manager.DeleteSecretRequest{SecretID: id, Region: scw.RegionNlAms}); err != nil {
		return err
	}
	return nil
}

func (scalewayWrapper *ScalewayWrapper) DeleteSecretVersion(id string, revision string) error {
	if _, err := scalewayWrapper.Api.DestroySecretVersion(&secret_manager.DestroySecretVersionRequest{SecretID: id, Region: scw.RegionNlAms, Revision: revision}); err != nil {
		return err
	}
	return nil
}
