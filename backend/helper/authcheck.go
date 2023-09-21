package decentproof_functions

import (
	"net/http"
	"os"
)

// Not yet functional
func authenticate(req *http.Request) bool {
	if isDebug, success := os.LookupEnv("DEBUG"); success {
		if isDebug == "true" {
			return true
		} else {
			///TODO: Implement this
			return verifyApiKey(req)

		}

	}
	return verifyApiKey(req)
}

func verifyApiKey(req *http.Request) bool {
	if req.Header.Get("Authorization") != "" {
		return true
	} else {
		return false
	}
}
