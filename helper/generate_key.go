package helper

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateApiKey(size int) string {
	bytes := make([]byte, size) // 32 bytes = 256 bits = 2^256 possible keys
	if _, err := rand.Read(bytes); err != nil {
		panic(err)
	} else {
		return base64.StdEncoding.EncodeToString(bytes)
	}
}
