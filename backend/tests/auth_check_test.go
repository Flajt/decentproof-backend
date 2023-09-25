package decentproof_functions

import (
	"net/http"
	"testing"

	helper "github.com/Flajt/decentproof-backend/decentproof-functions/helper"
)

func TestAuthCheck(t *testing.T) {
	t.Run("with valid api key", func(t *testing.T) {
		var apiKeys = []string{"test", "test2"}
		request := http.Request{Header: http.Header{"Authorization": []string{"Bearer " + apiKeys[0]}}, Method: "POST"}
		success := helper.VerifyApiKey(&request, apiKeys)
		if success != true {
			t.Errorf("Expected true, got %v", success)
		}
	})
	t.Run("with invalid api key", func(t *testing.T) {
		var apiKeys = []string{"test", "test2"}
		request := http.Request{Header: http.Header{"Authorization": []string{"Bearer " + "invalid"}}, Method: "POST"}
		success := helper.VerifyApiKey(&request, apiKeys)
		if success != false {
			t.Errorf("Expected false, got %v", success)
		}
	})
	t.Run("with no api key", func(t *testing.T) {
		var apiKeys = []string{"test", "test2"}
		request := http.Request{Header: http.Header{"Authorization": []string{""}}, Method: "POST"}
		success := helper.VerifyApiKey(&request, apiKeys)
		if success != false {
			t.Errorf("Expected false, got %v", success)
		}
	})

}
