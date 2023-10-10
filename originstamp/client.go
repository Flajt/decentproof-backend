package originstamp

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	models "github.com/Flajt/decentproof-backend/originstamp/models"
)

// Contains minimal version of the Originstamp API, not all api calls are included
// For a complete list of api calls see https://api.originstamp.com/swagger/swagger-ui.html
type OriginStampApiClient struct {
	ApiKey  string
	BaseUrl string
}

// Create Timestamping Client
func NewOriginStampApiClient(apiKey string) *OriginStampApiClient {
	return &OriginStampApiClient{
		ApiKey:  apiKey,
		BaseUrl: "https://api.originstamp.com/",
	}
}

// Create Timestamp
func (client *OriginStampApiClient) CreateTimestamp(body models.OriginStampProofRequestBody) (models.OriginStampCreateTimestampResponse, error) {
	encodedBody, err := json.Marshal(body)
	if err != nil {
		return models.OriginStampCreateTimestampResponse{}, err
	}
	reader := bytes.NewReader(encodedBody)
	resp, err := http.Post(client.BaseUrl+"v4/timestamp/create", "application/json", reader)
	if err != nil {
		return models.OriginStampCreateTimestampResponse{}, err
	}
	defer resp.Body.Close()

	var respBody models.OriginStampCreateTimestampResponse
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.OriginStampCreateTimestampResponse{}, err
	}
	json.Unmarshal(bodyBytes, &respBody)
	return respBody, nil

}

func (client *OriginStampApiClient) GetProof(body models.OriginStampProofRequestBody) (models.OriginStampProofResponse, error) {
	encodedBody, err := json.Marshal(body)
	if err != nil {
		return models.OriginStampProofResponse{}, err
	}
	reader := bytes.NewReader(encodedBody)

	resp, err := http.Post(client.BaseUrl+"v3/timestamp/proof/url", "application/json", reader)
	if err != nil {
		return models.OriginStampProofResponse{}, err
	}
	defer resp.Body.Close()

	var respBody models.OriginStampProofResponse
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.OriginStampProofResponse{}, err
	}
	json.Unmarshal(bodyBytes, &respBody)
	return respBody, nil
}

// Get Timestamp status by hash

func (client *OriginStampApiClient) GetTimestampStatus(hash string) (models.OriginStampTimeStampStatusResponse, error) {
	resp, err := http.Get(client.BaseUrl + "v4/timestamp/" + hash)
	if err != nil {
		return models.OriginStampTimeStampStatusResponse{}, err
	}
	defer resp.Body.Close()

	var respBody models.OriginStampTimeStampStatusResponse
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.OriginStampTimeStampStatusResponse{}, err
	}
	json.Unmarshal(bodyBytes, &respBody)
	return respBody, nil
}
