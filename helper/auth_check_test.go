package helper

import (
	"net/http"
	"testing"
)

func TestAuthCheck(t *testing.T) {
	t.Run("with valid api key", func(t *testing.T) {
		var apiKeys = []string{"test", "test2"}
		request := http.Request{Header: http.Header{"Authorization": []string{"Bearer " + apiKeys[0]}}, Method: "POST"}
		success, err := Authenticate(&request, apiKeys, false)
		if success != true || err != nil {
			t.Errorf("Expected true, got %v", success)
		}
	})
	t.Run("with invalid api key", func(t *testing.T) {
		var apiKeys = []string{"test", "test2"}
		request := http.Request{Header: http.Header{"Authorization": []string{"Bearer " + "invalid"}}, Method: "POST"}
		success, err := Authenticate(&request, apiKeys, false)
		if success != false || err != nil {
			t.Errorf("Expected false, got %v", success)
		}
	})
	t.Run("with no api key", func(t *testing.T) {
		var apiKeys = []string{"test", "test2"}
		request := http.Request{Header: http.Header{"Authorization": []string{""}}, Method: "POST"}
		success, err := Authenticate(&request, apiKeys, false)
		if success != false || err == nil {
			t.Errorf("Expected false, got %v", success)
		}
	})

}
