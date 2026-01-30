package server

import (
	"informatik/api/internal/store"
	"net/http"
)

type Server struct {
	store  store.Storer
	router http.Handler
}

func New(store store.Storer) *Server {
	s := &Server{
		store: store,
	}

	s.routes()
	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
