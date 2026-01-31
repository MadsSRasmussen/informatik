package server

import (
	"informatik/api/internal/ai"
	"informatik/api/internal/store"
	"net/http"
)

type Server struct {
	store    store.Storer
	aiClient ai.AIClient
	router   http.Handler
}

func New(store store.Storer, aiClient ai.AIClient) *Server {
	s := &Server{
		store:    store,
		aiClient: aiClient,
	}

	s.routes()
	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.corsMiddleware(s.router).ServeHTTP(w, r)
}
