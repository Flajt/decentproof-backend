package encryption_service

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
	"sort"
	"strconv"

	scw_secret_manager "github.com/Flajt/decentproof-backend/scw_secret_wrapper"
)

type EncryptionOutPut struct {
	Data  []byte
	Nonce []byte
}

type IEncryptionService interface {
	EncryptData(data []byte) (EncryptionOutPut, error)
	DecryptData(data []byte, nonce []byte) ([]byte, error)
}

type EncryptionService struct {
	scwWrapper *scw_secret_manager.ScalewayWrapper
	// Key for offline testing only, this can be used to replace the scwWrapper, 32bits are required!
	key *[]byte
}

// EncryptData encrypts the data using AES-GCM, the nonce is generated using the crypto/rand package
func (e *EncryptionService) EncryptData(data []byte) (EncryptionOutPut, error) {
	var block cipher.Block // To reduce code duplication, the if and else statements will asign their aes cipher to this variable
	if e.key != nil {
		blockCipher, err := aes.NewCipher(*e.key)
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
		sort.Slice(secretVersions.SecretVersions, func(i, j int) bool {
			return secretVersions.SecretVersions[i].CreatedAt.Before(*secretVersions.SecretVersions[j].CreatedAt)
		})
		secretVersionData := secretVersions.SecretVersions
		key, err := e.scwWrapper.GetSecretData(secret.Name, strconv.FormatUint(uint64(secretVersionData[0].Revision), 10))
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
	var secretVersionHolder scw_secret_manager.SecretVersionHolder
	var sharedSecrets scw_secret_manager.SecretHolder
	if e.key != nil {
		cipher, err := aes.NewCipher(*e.key)
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
		sort.Slice(secretVersions.SecretVersions, func(i, j int) bool {
			return secretVersions.SecretVersions[i].CreatedAt.Before(*secretVersions.SecretVersions[j].CreatedAt)
		})
		secretVersionData := secretVersions.SecretVersions
		key, err := e.scwWrapper.GetSecretData(secret.Name, strconv.FormatUint(uint64(secretVersionData[0].Revision), 10))
		if err != nil {
			return nil, err
		}
		cipher, err := aes.NewCipher(key)
		if err != nil {
			return nil, err
		}
		block = cipher
		secretVersionHolder = secretVersions
		sharedSecrets = secrets
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	outPut, err := aesgcm.Open(nil, nonce, data, nil)

	if err != nil {
		if err.Error() == "cipher: message authentication failed" && secretVersionHolder.TotalCount == 2 {
			secrertVersion := secretVersionHolder.SecretVersions[1]
			secret := sharedSecrets.Secrets[0]
			key2, err := e.scwWrapper.GetSecretData(secret.Name, strconv.FormatUint(uint64(secrertVersion.Revision), 10))
			if err != nil {
				return nil, err
			}
			block2, err := aes.NewCipher(key2)
			if err != nil {
				return nil, err
			}
			aesgcm2, err := cipher.NewGCM(block2)
			if err != nil {
				return nil, err
			}
			outPut, err := aesgcm2.Open(nil, nonce, data, nil)
			if err != nil {
				return nil, err
			}
			return outPut, nil
		} else {
			panic(err.Error())
		}
	}
	return outPut, nil
}

func NewEncryptionService(scwWrapper scw_secret_manager.ScalewayWrapper) IEncryptionService {
	return &EncryptionService{scwWrapper: &scwWrapper, key: nil}
}
func NewEncryptionServiceFromKey(keyAsBytes []byte) IEncryptionService {
	if len(keyAsBytes) != 32 {
		panic("Key is not 32 bytes long")
	}
	return &EncryptionService{scwWrapper: nil, key: &keyAsBytes}
}
