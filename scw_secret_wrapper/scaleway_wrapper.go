package scw_secret_manager

import (
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

type ScaleWaySetupData struct {
	AccessKey string
	SecretKey string
	ProjectID string
	Region    string
}

// Used godotenv to read you enviroment variables
func NewScaleWayWrapper(setupData ScaleWaySetupData) (*ScalewayWrapper, error) {

	if client, err := scw.NewClient(
		scw.WithAuth(setupData.AccessKey, setupData.SecretKey),
		scw.WithDefaultRegion(scw.Region(setupData.Region)),
		scw.WithDefaultProjectID(setupData.ProjectID),
	); err != nil {
		panic(err)
	} else {
		api := secret_manager.NewAPI(client)
		return &ScalewayWrapper{Client: *client, Api: api, PROJECT_ID: setupData.ProjectID}, nil
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

	if secret, err := scalewayWrapper.Api.CreateSecret(&secret_manager.CreateSecretRequest{Name: secretName, Type: secret_manager.SecretTypeUnknownSecretType}); err != nil {
		return err
	} else {
		if _, err := scalewayWrapper.Api.CreateSecretVersion(&secret_manager.CreateSecretVersionRequest{SecretID: secret.ID, Data: inputBytes}); err != nil {
			return err
		}
		return nil
	}
}

func (scalewayWrapper *ScalewayWrapper) CreateNewSecretVersion(secret secret_manager.Secret, data string) error {
	inputBytes := []byte(data)
	if _, err := scalewayWrapper.Api.CreateSecretVersion(&secret_manager.CreateSecretVersionRequest{SecretID: secret.ID, Data: inputBytes}); err != nil {
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
