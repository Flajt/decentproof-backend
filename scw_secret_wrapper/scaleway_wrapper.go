package scw_secret_manager

import (
	"os"

	secret_manager "github.com/scaleway/scaleway-sdk-go/api/secret/v1alpha1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

type Secret secret_manager.Secret

type SecretVersion secret_manager.SecretVersion

type SecretHolder struct {
	Secrets    []*Secret
	TotalCount uint32
}

type SecretVersionHolder struct {
	SecretVersions []SecretVersion
	TotalCount     uint32
}
type IScaleWayWrapper interface {
	ListSecrets(names ...string) (SecretHolder, error)
	ListSecretVersions(secretID string) (SecretVersionHolder, error)
	GetSecretData(secretName string, revision string) ([]byte, error)
	SetSecret(secretName string, secretValue []byte) (Secret, error)
	CreateNewSecretVersion(secret Secret, data []byte) error
	DeleteSecret(id string) error
	DeleteSecretVersion(id string, revision string) error
}

type ScalewayWrapper struct {
	Client     scw.Client
	Api        *secret_manager.API
	PROJECT_ID string
}

type ScaleWaySetupData struct {
	AccessKey string
	SecretKey string
	ProjectID string
	Region    string
}

// Used godotenv to read you enviroment variables
func NewScaleWayWrapper(setupData ScaleWaySetupData) IScaleWayWrapper {

	if client, err := scw.NewClient(
		scw.WithAuth(setupData.AccessKey, setupData.SecretKey),
		scw.WithDefaultRegion(scw.Region(setupData.Region)),
		scw.WithDefaultProjectID(setupData.ProjectID),
	); err != nil {
		panic(err)
	} else {
		api := secret_manager.NewAPI(client)
		return &ScalewayWrapper{Client: *client, Api: api, PROJECT_ID: setupData.ProjectID}
	}
}
func NewScaleWayWrapperFromEnv() IScaleWayWrapper {
	return NewScaleWayWrapper(ScaleWaySetupData{AccessKey: os.Getenv("SCW_ACCESS_KEY"), SecretKey: os.Getenv("SCW_SECRET_KEY"), ProjectID: os.Getenv("SCW_DEFAULT_PROJECT_ID"), Region: os.Getenv("SCW_DEFAULT_REGION")})
}

func (scalewayWrapper *ScalewayWrapper) ListSecrets(names ...string) (SecretHolder, error) {
	if len(names) > 0 {
		if secrets, err := scalewayWrapper.Api.ListSecrets(&secret_manager.ListSecretsRequest{ProjectID: &scalewayWrapper.PROJECT_ID, Name: &names[0]}); err != nil {
			return SecretHolder{}, err
		} else {
			secretStore := make([]*Secret, len(secrets.Secrets))
			for i, secret := range secrets.Secrets {
				secretStore[i] = (*Secret)(secret)
			}
			return SecretHolder{Secrets: secretStore, TotalCount: secrets.TotalCount}, nil
		}
	} else {
		if secrets, err := scalewayWrapper.Api.ListSecrets(&secret_manager.ListSecretsRequest{ProjectID: &scalewayWrapper.PROJECT_ID}); err != nil {
			return SecretHolder{}, err
		} else {
			secretStore := make([]*Secret, len(secrets.Secrets))
			for i, secret := range secrets.Secrets {
				secretStore[i] = (*Secret)(secret)
			}
			return SecretHolder{Secrets: secretStore, TotalCount: secrets.TotalCount}, nil
		}
	}

}

func (ScalewayWrapper *ScalewayWrapper) ListSecretVersions(secretID string) (SecretVersionHolder, error) {
	if secrets, err := ScalewayWrapper.Api.ListSecretVersions(&secret_manager.ListSecretVersionsRequest{SecretID: secretID}); err != nil {
		return SecretVersionHolder{}, err
	} else {
		secretVersions := make([]SecretVersion, len(secrets.Versions))
		for i, secretVersion := range secrets.Versions {
			secretVersions[i] = SecretVersion(*secretVersion)
		}
		return SecretVersionHolder{SecretVersions: secretVersions, TotalCount: secrets.TotalCount}, nil
	}
}

func (scalewayWrapper *ScalewayWrapper) GetSecretData(secretName string, revision string) ([]byte, error) {
	requestParams := &secret_manager.GetSecretVersionByNameRequest{Region: scw.RegionNlAms, Revision: revision, SecretName: secretName, ProjectID: &scalewayWrapper.PROJECT_ID}
	if secret, err := scalewayWrapper.Api.GetSecretVersionByName(requestParams); err != nil {
		return []byte{}, err
	} else {
		if secretVersion, err := scalewayWrapper.Api.AccessSecretVersion(&secret_manager.AccessSecretVersionRequest{Region: scw.RegionNlAms, SecretID: secret.SecretID, Revision: revision}); err != nil {
			return []byte{}, err
		} else {
			return secretVersion.Data, nil
		}
	}
}

func (scalewayWrapper *ScalewayWrapper) SetSecret(secretName string, secretValue []byte) (Secret, error) {

	secret, err := scalewayWrapper.Api.CreateSecret(&secret_manager.CreateSecretRequest{Name: secretName, Type: secret_manager.SecretTypeUnknownSecretType})
	if err != nil {
		return Secret{}, err
	}
	if _, err := scalewayWrapper.Api.CreateSecretVersion(&secret_manager.CreateSecretVersionRequest{SecretID: secret.ID, Data: secretValue}); err != nil {
		return Secret{}, err
	}
	return Secret(*secret), nil

}

func (scalewayWrapper *ScalewayWrapper) CreateNewSecretVersion(secret Secret, data []byte) error {
	if _, err := scalewayWrapper.Api.CreateSecretVersion(&secret_manager.CreateSecretVersionRequest{SecretID: secret.ID, Data: data}); err != nil {
		return err
	}
	return nil
}

func (scalewayWrapper *ScalewayWrapper) DeleteSecret(id string) error {
	if err := scalewayWrapper.Api.DeleteSecret(&secret_manager.DeleteSecretRequest{SecretID: id}); err != nil {
		return err
	}
	return nil
}

func (scalewayWrapper *ScalewayWrapper) DeleteSecretVersion(id string, revision string) error {
	if _, err := scalewayWrapper.Api.DestroySecretVersion(&secret_manager.DestroySecretVersionRequest{SecretID: id, Revision: revision}); err != nil {
		return err
	}
	return nil
}
