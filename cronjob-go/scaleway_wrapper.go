package decentproof_cronjob

import (
	secret_manager "github.com/scaleway/scaleway-sdk-go/api/secret/v1alpha1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

type SecretHolder struct {
	Secrets    []*secret_manager.Secret
	totalCount uint32
}

type SecretVersionHolder struct {
	SecretVersions []*secret_manager.SecretVersion
	totalCount     uint32
}

type ScalewayWrapper struct {
	Client scw.Client
	Api    *secret_manager.API
}

func NewAppcheckWrapper() (*ScalewayWrapper, error) {
	if client, err := scw.NewClient(
		scw.WithEnv(),
	); err != nil {
		return nil, err
	} else {
		api := secret_manager.NewAPI(client)
		return &ScalewayWrapper{Client: *client, Api: api}, nil
	}
}

func (scalewayWrapper *ScalewayWrapper) ListSecrets() (SecretHolder, error) {
	if secrets, err := scalewayWrapper.Api.ListSecrets(&secret_manager.ListSecretsRequest{}); err != nil {
		return SecretHolder{}, err
	} else {
		return SecretHolder{Secrets: secrets.Secrets, totalCount: secrets.TotalCount}, nil
	}
}

func (ScalewayWrapper *ScalewayWrapper) ListSecretVersions(secretID string) (SecretVersionHolder, error) {
	if secrets, err := ScalewayWrapper.Api.ListSecretVersions(&secret_manager.ListSecretVersionsRequest{SecretID: secretID}); err != nil {
		return SecretVersionHolder{}, err
	} else {
		return SecretVersionHolder{SecretVersions: secrets.Versions, totalCount: secrets.TotalCount}, nil
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

	if secret, err := scalewayWrapper.Api.CreateSecret(&secret_manager.CreateSecretRequest{ProjectID: *scw.LoadEnvProfile().DefaultProjectID, Region: scw.RegionNlAms, Name: secretName, Type: secret_manager.SecretTypeUnknownSecretType}); err != nil {
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
