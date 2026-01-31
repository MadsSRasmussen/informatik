package server

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRespond(t *testing.T) {
	t.Run("json", func(t *testing.T) {
		server := &Server{}
		rec := httptest.NewRecorder()

		payload := map[string]string{"foo": "bar"}
		server.respond(rec, http.StatusOK, payload)

		if rec.Code != http.StatusOK {
			t.Fatalf("expected status %d, got %d", http.StatusOK, rec.Code)
		}

		var out map[string]string
		if err := json.NewDecoder(rec.Body).Decode(&out); err != nil {
			t.Fatal(err)
		}
	})
}
