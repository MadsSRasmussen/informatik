package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"informatik/api/internal/store"
	"net/http"
	"strconv"
)

func (s *Server) handlePostsList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, err := s.store.ListPosts()
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		s.respond(w, http.StatusOK, res)
	}
}

func (s *Server) handlePostGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		res, err := s.store.GetPost(id)
		if err != nil {
			if errors.Is(err, store.ErrNotFound) {
				http.Error(
					w,
					fmt.Sprintf("Post with id %d not found", id),
					http.StatusNotFound,
				)
				return
			}

			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		s.respond(w, http.StatusOK, res)
	}
}

type PostCreateRequest struct {
	Content string `json:"content"`
}

func (s *Server) handlePostCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req PostCreateRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		res, err := s.store.CreatePost(req.Content)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		s.respond(w, http.StatusCreated, res)
	}
}
