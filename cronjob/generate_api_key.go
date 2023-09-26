package decentproof_cronjob

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateApiKey() string {
	bytes := make([]byte, 32) // 32 bytes = 256 bits = 2^256 possible keys
	if _, err := rand.Read(bytes); err != nil {
		panic(err)
	} else {
		return base64.StdEncoding.EncodeToString(bytes)
	}

}
