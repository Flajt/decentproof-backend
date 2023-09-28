package sign

import (
	"encoding/json"
	"net/http"

	"github.com/Flajt/decentproof-backend/helper"
	secret_wrapper "github.com/Flajt/decentproof-backend/scw_secret_wrapper"
)

func HandleSignature(w http.ResponseWriter, r *http.Request) {
	isValid := helper.VerifyApiKey(r, helper.RetrievApiKeys())
	if !isValid {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized"))
		return
	}

	scw_wrapper := secret_wrapper.NewScaleWayWrapperFromEnv()
	signatureManager := NewSignatureManager(scw_wrapper)
	signatureManager.InitSignatureManager()
	jsonDecoder := json.NewDecoder(r.Body)
	var holder SignatureRequestBody
	if err := jsonDecoder.Decode(&holder); err != nil {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request"))
		return
	}
	signatureBytes, err := signatureManager.SignData([]byte(holder.Data))
	if err != nil {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}
	signatureResp := SignatureResponseBody{Signature: string(signatureBytes)}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	respBytes, err := json.Marshal(signatureResp)
	if err != nil {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}
	w.Write(respBytes)
}
