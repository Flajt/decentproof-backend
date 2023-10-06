package verify_hash

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/Flajt/decentproof-backend/helper"
	"github.com/Flajt/decentproof-backend/originstamp"
	models "github.com/Flajt/decentproof-backend/verify-hash/model"
)

func HandleHashVerification(w http.ResponseWriter, r *http.Request) {
	apiKeys := helper.RetrievApiKeys()
	isValid := helper.VerifyApiKey(r, apiKeys)
	if !isValid {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized"))
		return
	}
	requestModel := models.VerifyRequestBody{}
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	defer r.Body.Close()
	err = json.Unmarshal(bytes, &requestModel)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	if requestModel.Hash == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Hash is required"))
		return
	}
	ORIGINSTAMP_API_KEY := os.Getenv("ORIGINSTAMP_API_KEY")
	orignstampClient := originstamp.NewOriginStampApiClient(ORIGINSTAMP_API_KEY)
	resp, err := orignstampClient.GetTimestampStatus(requestModel.Hash)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	if resp.ErrorMessage != "" {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(resp.ErrorMessage))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
