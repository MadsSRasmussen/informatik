package server

import (
	"encoding/json"
	"fmt"
	"informatik/api/internal/ai"
	"net/http"
)

func (s *Server) handleCompletionCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req ai.CompletionsRequestBody
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid json", http.StatusBadRequest)
			return
		}

		res, err := s.aiClient.GetCompletion(&req)
		if err != nil {
			fmt.Printf("Error occured: %v\n", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		s.respond(w, http.StatusOK, res)
	}
}
