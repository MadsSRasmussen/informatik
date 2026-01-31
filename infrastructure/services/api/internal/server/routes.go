package server

import "net/http"

func (s *Server) routes() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /ping", s.handlePing())

	mux.HandleFunc("GET /posts", s.handlePostsList())
	mux.HandleFunc("GET /posts/{id}", s.handlePostGet())
	mux.HandleFunc("POST /posts", s.handlePostCreate())

	mux.HandleFunc("POST /completions", s.handleCompletionCreate())

	s.router = mux
}
