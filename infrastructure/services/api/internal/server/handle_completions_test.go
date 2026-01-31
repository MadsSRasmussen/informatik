package server

import (
	"errors"
	"informatik/api/internal/ai"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandleCompletionCreate(t *testing.T) {
	tests := []struct {
		name           string
		body           string
		clientError    error
		expectedStatus int
	}{
		{
			name:           "success",
			body:           `{"model":"mistral-medium-latest","messages":[{"role":"user","content":"hi"}]}`,
			clientError:    nil,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "invalid_json",
			body:           `{"model":"mistral-medium-latest","messages":`,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "client_error",
			body:           `{"model":"mistral-medium-latest","messages":[{"role":"user","content":"hi"}]}`,
			clientError:    errors.New("error"),
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &MockAIClient{
				GetCompletionFunc: func(r *ai.CompletionsRequestBody) (*ai.Message, error) {
					return &ai.Message{
						Role:    ai.RoleAssistant,
						Content: "Mock response",
					}, tt.clientError
				},
			}

			server := New(&MockStore{}, client)
			req := httptest.NewRequest(
				http.MethodPost,
				"/completions",
				strings.NewReader(tt.body),
			)
			rec := httptest.NewRecorder()

			server.ServeHTTP(rec, req)

			if rec.Code != tt.expectedStatus {
				t.Fatalf("expected %d, got %d", tt.expectedStatus, rec.Code)
			}
		})
	}
}
