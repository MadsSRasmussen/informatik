package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlePing(t *testing.T) {
	tests := []struct {
		name           string
		expectedStatus int
	}{
		{"success", http.StatusOK},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := New(&MockStore{}, &MockAIClient{})
			req := httptest.NewRequest(http.MethodGet, "/ping", nil)
			rec := httptest.NewRecorder()

			server.ServeHTTP(rec, req)

			if rec.Code != tt.expectedStatus {
				t.Fatalf("expected %d, got %d", tt.expectedStatus, rec.Code)
			}
		})
	}
}
