package decentproof_functions

import (
	"net/http"

	decentproof_functions "github.com/Flajt/decentproof-backend/decentproof-functions/helper"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("X-Appcheck")
	if authHeader == "" {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized"))
		return
	}
	appCheckWrapper := decentproof_functions.NewAppcheckWrapper()
	success, err := appCheckWrapper.CheckApp(authHeader)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized"))
		return
	} else {
		apiKeys := decentproof_functions.RetrievApiKeys()
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
