package decentproof_functions

import (
	"encoding/json"
	"net/http"

	helper "github.com/Flajt/decentproof-backend/decentproof-functions/helper"
)

func HandleHasNewKey(w http.ResponseWriter, r *http.Request) {
	keys := helper.RetrievApiKeys()
	requestKey := r.Header.Get("Authorization")
	isValid := helper.VerifyApiKey(r, keys)
	if isValid {
		if keys[0] == requestKey {
			response := map[string]bool{"hasNewKey": true}
			responseBytes, err := json.Marshal(response)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("Internal Server Error"))
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			w.Write(responseBytes)
			return
		}

	} else {
		w.WriteHeader(http.StatusUnauthorized)
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("Unauthorized"))
		return
	}
}
