package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (s *Server) respond(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if data != nil {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			fmt.Printf("Error encoding response: %v\n", err)
		}
	}
}
