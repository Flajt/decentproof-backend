package encryption_service

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
	"strconv"

	scw_secret_manager "github.com/Flajt/decentproof-backend/scw_secret_wrapper"
)

type EncryptionOutPut struct {
	Data  []byte
	Nonce []byte
}

type IEncryptionService interface {
	EncryptData(data []byte) (EncryptionOutPut, error)
	DecryptData(data []byte, nonce []byte) (EncryptionOutPut, error)
	NewEncryptionService(scwWrapper scw_secret_manager.ScalewayWrapper) IEncryptionService
	NewEncryptionServiceFromEnv() IEncryptionService
}

type EncryptionService struct {
	scwWrapper *scw_secret_manager.ScalewayWrapper
	// Key for offline testing only, this can be used to replace the scwWrapper, 32bits are required!
	key *string
}

// EncryptData encrypts the data using AES-GCM, the nonce is generated using the crypto/rand package
func (e *EncryptionService) EncryptData(data []byte) (EncryptionOutPut, error) {
	var block cipher.Block // To reduce code duplication, the if and else statements will asign their aes cipher to this variable
	if e.key != nil {
		blockCipher, err := aes.NewCipher([]byte(*e.key))
		if err != nil {
			return EncryptionOutPut{}, err
		}
		block = blockCipher
	} else {
		secrets, err := e.scwWrapper.ListSecrets("ENCRYPTION_KEY")
		if err != nil {
			return EncryptionOutPut{}, err
		}
		if secrets.TotalCount == 0 {
			panic("No encryption key found")
		}
		secret := secrets.Secrets[0]
		secretVersions, err := e.scwWrapper.ListSecretVersions(secret.ID)
		if err != nil {
			return EncryptionOutPut{}, err
		}
		var secretVersionIndex = -2
		for index, secretVersion := range secretVersions.SecretVersions {
			if secretVersion.IsLatest {
				secretVersionIndex = index
				break
			}
		}
		if secretVersionIndex == -2 {
			panic("Secret not found")
		}
		secretVersionData := secretVersions.SecretVersions[secretVersionIndex]
		key, err := e.scwWrapper.GetSecretData(secret.Name, strconv.FormatUint(uint64(secretVersionData.Revision), 10))
		if err != nil {
			return EncryptionOutPut{}, err
		}
		blockCipher, err := aes.NewCipher(key)
		if err != nil {
			return EncryptionOutPut{}, err
		}
		block = blockCipher
	}
	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return EncryptionOutPut{}, err
	}
	outPut := aesgcm.Seal(nil, nonce, data, nil)
	if len(outPut) != 0 {
		return EncryptionOutPut{Nonce: nonce, Data: outPut}, nil
	} else {
		return EncryptionOutPut{}, errors.New("empty buffer, no data to encrypt")
	}
}

// DecryptData decrypts the data using AES-GCM, the nonce needs to be provided
func (e *EncryptionService) DecryptData(data []byte, nonce []byte) ([]byte, error) {
	var block cipher.Block // Same as the function above
	if e.key != nil {
		cipher, err := aes.NewCipher([]byte(*e.key))
		if err != nil {
			return nil, err
		}
		block = cipher
	} else {
		secrets, err := e.scwWrapper.ListSecrets("ENCRYPTION_KEY")
		if err != nil {
			return nil, err
		}
		if secrets.TotalCount == 0 {
			panic("No encryption key found")
		}
		secret := secrets.Secrets[0]
		secretVersions, err := e.scwWrapper.ListSecretVersions(secret.ID)
		if err != nil {
			return nil, err
		}
		var secretVersionIndex = -2
		for index, secretVersion := range secretVersions.SecretVersions {
			if secretVersion.IsLatest {
				secretVersionIndex = index
				break
			}
		}
		if secretVersionIndex == -2 {
			panic("Secret not found")
		}
		secretVersionData := secretVersions.SecretVersions[secretVersionIndex]
		key, err := e.scwWrapper.GetSecretData(secret.Name, strconv.FormatUint(uint64(secretVersionData.Revision), 10))
		if err != nil {
			return nil, err
		}
		cipher, err := aes.NewCipher(key)
		if err != nil {
			return nil, err
		}
		block = cipher
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	outPut, err := aesgcm.Open(nil, nonce, data, nil)
	if err != nil {
		panic(err.Error())
	}
	return outPut, nil
}

func (e *EncryptionService) NewEncryptionService(scwWrapper scw_secret_manager.ScalewayWrapper) EncryptionService {
	return EncryptionService{scwWrapper: &scwWrapper, key: nil}
}

func (e *EncryptionService) NewEncryptionServiceFromEnv(key string) EncryptionService {
	return EncryptionService{scwWrapper: nil, key: &key}
}
