package decentproof_cronjob

import (
	"net/http"
	"testing"

	decentproof_cronjob "github.com/decentproof-cron"
)

// MockResponseWriter is a basic mock implementation of http.ResponseWriter.
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

func TestCronjob(t *testing.T) {
	t.Run("with zero entrie", func(t *testing.T) {
		wrapper, _ := decentproof_cronjob.NewScaleWayWrapper()
		decentproof_cronjob.Handle(&MockResponseWriter{}, nil)
		want := 1
		secretHolder, err := wrapper.ListSecrets()
		if err != nil {
			t.Error(err)
		}
		if secretHolder.TotalCount != uint32(want) {
			t.Errorf("Got %d secrets, wanted %d", secretHolder.TotalCount, want)
		}
		versionHolder, err := wrapper.ListSecretVersions(secretHolder.Secrets[0].ID)
		if err != nil {
			t.Error(err)
		}
		if versionHolder.TotalCount != uint32(want) {
			t.Errorf("Got %d secrets, wanted %d", versionHolder.TotalCount, want)
		}
		cleanUp(t)
	})

	t.Run("with two entries", func(t *testing.T) {
		wrapper, _ := decentproof_cronjob.NewScaleWayWrapper()
		if err := wrapper.SetSecret("apiKey", "test"); err != nil {
			t.Error(err)
		}
		secretHolder1, err := wrapper.ListSecrets("apiKey")
		if err != nil {
			t.Error(err)
		}
		if err := wrapper.CreateNewSecretVersion(*secretHolder1.Secrets[0], "test2"); err != nil {
			t.Error(err)
		}
		decentproof_cronjob.Handle(&MockResponseWriter{}, nil)
		secretWant := 1
		versionWant := 2
		secretHolder, err := wrapper.ListSecrets()
		if err != nil {
			t.Error(err)
		}
		if secretHolder.TotalCount != uint32(secretWant) {
			t.Errorf("Got %d secrets, wanted %d", secretHolder.TotalCount, secretWant)
		}
		versionHolder, err := wrapper.ListSecretVersions(secretHolder.Secrets[0].ID)
		if err != nil {
			t.Error(err)
		}
		if versionHolder.TotalCount != uint32(versionWant) {
			t.Errorf("Got %d secrets, wanted %d", versionHolder.TotalCount, versionWant)
		}
	})

	t.Cleanup(func() { cleanUp(t) })
}
