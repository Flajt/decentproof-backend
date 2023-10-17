package originstamp

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"

	models "github.com/Flajt/decentproof-backend/originstamp/models"
)

// Contains minimal version of the Originstamp API, not all api calls are included
// For a complete list of api calls see https://api.originstamp.com/swagger/swagger-ui.html
type OriginStampApiClient struct {
	ApiKey     string
	BaseUrl    string
	httpClient http.Client
}

// Create Timestamping Client
func NewOriginStampApiClient(apiKey string) *OriginStampApiClient {
	return &OriginStampApiClient{
		ApiKey:     apiKey,
		BaseUrl:    "https://api.originstamp.com/",
		httpClient: http.Client{Timeout: 5 * time.Second},
	}
}

// Create Timestamp
func (client *OriginStampApiClient) CreateTimestamp(body models.OriginStampTimestampRequestBody) (models.OriginStampCreateTimestampResponse, error) {
	encodedBody, err := json.Marshal(body)
	if err != nil {
		return models.OriginStampCreateTimestampResponse{}, err
	}
	reader := bytes.NewReader(encodedBody)
	req, err := http.NewRequest(http.MethodPost, client.BaseUrl+"v4/timestamp/create", reader)
	if err != nil {
		return models.OriginStampCreateTimestampResponse{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", client.ApiKey)
	resp, err := client.httpClient.Do(req)
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

	req, err := http.NewRequest(http.MethodPost, client.BaseUrl+"v3/timestamp/proof/url", reader)
	if err != nil {
		return models.OriginStampProofResponse{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", client.ApiKey)
	resp, err := client.httpClient.Do(req)
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
	req, err := http.NewRequest(http.MethodGet, client.BaseUrl+"v4/timestamp/"+hash, nil)
	if err != nil {
		return models.OriginStampTimeStampStatusResponse{}, err
	}
	req.Header.Set("Authorization", client.ApiKey)
	resp, err := client.httpClient.Do(req)
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
