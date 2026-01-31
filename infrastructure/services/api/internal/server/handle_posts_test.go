package server

import (
	"errors"
	"fmt"
	"informatik/api/internal/db"
	"informatik/api/internal/store"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestHandlePostsList(t *testing.T) {
	tests := []struct {
		name           string
		storeError     error
		expectedStatus int
	}{
		{"success", nil, http.StatusOK},
		{"store error", errors.New("error"), http.StatusInternalServerError},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := &MockStore{
				ListPostsFunc: func() ([]db.Post, error) {
					return []db.Post{
						{ID: 1, Content: "Mock post", CreatedAt: time.Now()},
					}, tt.storeError
				},
			}

			server := New(store, &MockAIClient{})
			req := httptest.NewRequest(http.MethodGet, "/posts", nil)
			rec := httptest.NewRecorder()

			server.ServeHTTP(rec, req)

			if rec.Code != tt.expectedStatus {
				t.Fatalf("expected %d, got %d", tt.expectedStatus, rec.Code)
			}
		})
	}
}

func TestHandlePostGet(t *testing.T) {
	tests := []struct {
		name           string
		storeError     error
		postID         string
		expectedStatus int
	}{
		{"success", nil, "1", http.StatusOK},
		{"not found", store.ErrNotFound, "99", http.StatusNotFound},
		{"store error", errors.New("error"), "1", http.StatusInternalServerError},
		{"invalid id", nil, "abc", http.StatusInternalServerError},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := &MockStore{
				GetPostFunc: func(pid int) (*db.Post, error) {
					if pid == 1 {
						return &db.Post{
							ID:        1,
							Content:   "Mock post",
							CreatedAt: time.Now(),
						}, tt.storeError
					} else {
						return nil, store.ErrNotFound
					}
				},
			}

			server := New(store, &MockAIClient{})
			req := httptest.NewRequest(
				http.MethodGet,
				fmt.Sprintf("/posts/%s", tt.postID),
				nil,
			)
			rec := httptest.NewRecorder()

			server.ServeHTTP(rec, req)

			if rec.Code != tt.expectedStatus {
				t.Fatalf("expected %d, got %d", tt.expectedStatus, rec.Code)
			}
		})
	}
}

func TestHandlePostCreate(t *testing.T) {
	tests := []struct {
		name           string
		body           string
		storeError     error
		expectedStatus int
	}{
		{
			name:           "success",
			body:           `{"content":"mock"}`,
			storeError:     nil,
			expectedStatus: http.StatusCreated,
		},
		{
			name:           "invalid_json",
			body:           `{"content":`,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "store_error",
			body:           `{"content":"mock"}`,
			storeError:     errors.New("error"),
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := &MockStore{
				CreatePostFunc: func(content string) (*store.PostCreateResponse, error) {
					return &store.PostCreateResponse{ID: 1}, tt.storeError
				},
			}

			server := New(store, &MockAIClient{})
			req := httptest.NewRequest(
				http.MethodPost,
				"/posts",
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
