package discordlogger

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewDiscordLogger(t *testing.T) {
	dl := NewDiscordLogger("testurl", false)

	if dl.webhookUrl != "testurl" {
		t.Errorf("webhook url not what is expected, got: %v", dl.webhookUrl)
	}
}

func TestDiscordLogger_Log(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		bodyBytes, err := io.ReadAll(r.Body)
		if err != nil {
			t.Errorf("error reading log body: %v", err)
		}

		var reqBody discordWebhookPayload
		err = json.Unmarshal(bodyBytes, &reqBody)
		if err != nil {
			t.Errorf("error unmarshalling json: %v", err)
		}

		if reqBody.Content != "test log: 1" {
			t.Errorf("bad discord log message, got: %v", reqBody.Content)
		}

		w.WriteHeader(http.StatusNoContent)
		io.WriteString(w, "")
	}))
	defer mockServer.Close()

	testVar := 1
	dl := NewDiscordLogger(mockServer.URL, false)

	dl.Printf("test log: %v", testVar)
}

func TestDiscordLogger_Error(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		bodyBytes, err := io.ReadAll(r.Body)
		if err != nil {
			t.Errorf("error reading log error body: %v", err)
		}

		var reqBody discordWebhookPayload
		err = json.Unmarshal(bodyBytes, &reqBody)
		if err != nil {
			t.Errorf("error unmarshalling json: %v", err)
		}

		if reqBody.Content != "**Error:** test error: 2" {
			t.Errorf("bad discord error log message, got: %v", reqBody.Content)
		}

		w.WriteHeader(http.StatusNoContent)
		io.WriteString(w, "")
	}))
	defer mockServer.Close()

	testVar := 2
	dl := NewDiscordLogger(mockServer.URL, false)

	dl.Errorf("test error: %v", testVar)
}
