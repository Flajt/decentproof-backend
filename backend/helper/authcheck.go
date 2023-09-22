package decentproof_functions

import (
	"net/http"
	"strings"
)

func VerifyApiKey(req *http.Request, apiKeys []string) bool {
	if authHeader := req.Header.Get("Authorization"); authHeader != "" {
		apiKey := strings.Split(authHeader, " ")[1]
		match := false
		for _, key := range apiKeys {
			if apiKey == key {
				match = true
				break
			}
		}
		return match
	} else {
		return false
	}
}
