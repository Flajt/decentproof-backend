package has_new_key

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/Flajt/decentproof-backend/helper"
)

func HandleHasNewKey(w http.ResponseWriter, r *http.Request) {
	keys := helper.RetrievApiKeys()
	requestKey := strings.Split(r.Header.Get("Authorization"), " ")[1]
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
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(responseBytes)
			return
		} else if keys[1] == requestKey {
			response := map[string]bool{"hasNewKey": false}
			responseBytes, err := json.Marshal(response)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("Internal Server Error"))
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(responseBytes)
			return
		} else {
			response := map[string]bool{"hasNewKey": true}
			responseBytes, err := json.Marshal(response)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("Internal Server Error"))
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
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
