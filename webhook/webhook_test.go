package webhook

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"testing"

	encryption_service "github.com/Flajt/decentproof-backend/encryption"
	models "github.com/Flajt/decentproof-backend/originstamp/models"
	scw_secret_manager "github.com/Flajt/decentproof-backend/scw_secret_wrapper"
)

type MockResponseWriter struct {
	StatusCode int
	HeaderMap  http.Header
	Body       []byte
}

func (m *MockResponseWriter) WriteHeader(statusCode int) {
	m.StatusCode = statusCode
}

func (m *MockResponseWriter) Header() http.Header {
	if m.HeaderMap == nil {
		m.HeaderMap = make(http.Header)
	}
	return m.HeaderMap
}

func (m *MockResponseWriter) Write(data []byte) (int, error) {
	m.Body = append(m.Body, data...)
	return len(data), nil
}

func TestWebhookHandler(t *testing.T) {
	t.Run("Valid input + email", func(t *testing.T) {
		scwWrapper := scw_secret_manager.NewScaleWayWrapperForDev()
		encryptionService := encryption_service.NewEncryptionService(scwWrapper)
		validEmail := "myEmail@myDomain.com"
		encryptionData, err := encryptionService.EncryptData([]byte(validEmail))
		hexEmail := hex.EncodeToString(encryptionData.Data)
		hexNonce := hex.EncodeToString(encryptionData.Nonce)
		if err != nil {
			t.Errorf("Error encrypting email: %v", err)
		}
		reqBody := models.OriginStampWebhookRequestBody{
			Created:     false,
			DateCreated: 1541203188245,
			Comment:     "no comment",
			HashString:  "72c40efa8887c7ea583edf9a54ab0f2ea8cb92d394564ca83763104c2218d12a",
			Timestamps: []models.OriginStampTimeStamp{
				{
					SeedID:       "f98a61e8-22bd-4ea5-ab9e-02723ff78c40",
					CurrencyID:   0,
					Transaction:  "aed3db9ef94953f65e93d56a4e5bcf234d43e27a1b3e7ce0f274cc7ed750d0e2",
					PrivateKey:   "5e92ec09501a5d39e251a151f84b5e2228312c445eb23b4e1de6360e27bad54b",
					Timestamp:    1541203656000,
					SubmitStatus: 3,
				},
			},
		}
		reqBytes, err := json.Marshal(reqBody)
		if err != nil {
			t.Errorf("Error marshalling request body: %v", err)
		}

		req, err := http.NewRequest("POST", "?mail="+hexEmail+"&nonce="+hexNonce, bytes.NewBuffer(reqBytes))
		if err != nil {
			t.Errorf("Error creating request: %v", err)
		}
		writer := &MockResponseWriter{}
		HandleWebhookCallBack(writer, req)
		if writer.StatusCode != http.StatusOK {
			t.Errorf("Expected status code %v, got %v", http.StatusOK, writer.StatusCode)
		}
	})
	t.Run("Valid input + no email", func(t *testing.T) {
		emptyMail := ""
		reqBody := models.OriginStampWebhookRequestBody{
			Created:     false,
			DateCreated: 1541203188245,
			Comment:     "no comment",
			HashString:  "72c40efa8887c7ea583edf9a54ab0f2ea8cb92d394564ca83763104c2218d12a",
			Timestamps: []models.OriginStampTimeStamp{
				{
					SeedID:       "f98a61e8-22bd-4ea5-ab9e-02723ff78c40",
					CurrencyID:   0,
					Transaction:  "aed3db9ef94953f65e93d56a4e5bcf234d43e27a1b3e7ce0f274cc7ed750d0e2",
					PrivateKey:   "5e92ec09501a5d39e251a151f84b5e2228312c445eb23b4e1de6360e27bad54b",
					Timestamp:    1541203656000,
					SubmitStatus: 3,
				},
			},
		}
		reqBytes, err := json.Marshal(reqBody)
		if err != nil {
			t.Errorf("Error marshalling request body: %v", err)
		}

		req, err := http.NewRequest("POST", "/?mail="+emptyMail, bytes.NewBuffer(reqBytes))
		if err != nil {
			t.Errorf("Error creating request: %v", err)
		}
		writer := &MockResponseWriter{}
		HandleWebhookCallBack(writer, req)
		if writer.StatusCode != http.StatusBadRequest {
			t.Errorf("Expected status code %v, got %v", http.StatusOK, writer.StatusCode)
		}

	})
}
