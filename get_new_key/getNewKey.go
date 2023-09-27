package get_new_key

import (
	"net/http"

	"github.com/Flajt/decentproof-backend/helper"
)

func HandleGetNewKey(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("X-Appcheck")
	if authHeader == "" {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized"))
		return
	}
	appCheckWrapper := NewAppcheckWrapper()
	success, err := appCheckWrapper.CheckApp(authHeader)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized"))
		return
	} else {
		apiKeys := helper.RetrievApiKeys()
		if success {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(apiKeys[1]))
			return
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
			return
		}
	}
}
