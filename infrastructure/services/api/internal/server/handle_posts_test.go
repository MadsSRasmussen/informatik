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
		storeErr       error
		expectedStatus int
	}{
		{"success", nil, http.StatusOK},
		{"store_error", errors.New("error"), http.StatusInternalServerError},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := &MockStore{
				ListPostsFunc: func() ([]db.Post, error) {
					return []db.Post{
						{ID: 1, Content: "Mock post", CreatedAt: time.Now()},
					}, tt.storeErr
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
		storeErr       error
		postID         string
		expectedStatus int
	}{
		{"success", nil, "1", http.StatusOK},
		{"not_found", store.ErrNotFound, "99", http.StatusNotFound},
		{"store_error", errors.New("error"), "1", http.StatusInternalServerError},
		{"invalid_id", nil, "abc", http.StatusInternalServerError},
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
						}, tt.storeErr
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
		storeErr       error
		expectedStatus int
	}{
		{
			name:           "success",
			body:           `{"content":"mock"}`,
			storeErr:       nil,
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
			storeErr:       errors.New("error"),
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := &MockStore{
				CreatePostFunc: func(content string) (*store.PostCreateResponse, error) {
					return &store.PostCreateResponse{ID: 1}, tt.storeErr
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

func TestHandlePostRemove(t *testing.T) {
	tests := []struct {
		name           string
		storeErr       error
		postID         string
		expectedStatus int
	}{
		{
			name:           "success",
			storeErr:       nil,
			postID:         "1",
			expectedStatus: http.StatusNoContent,
		},
		{
			name:           "not found",
			storeErr:       store.ErrNotFound,
			postID:         "99",
			expectedStatus: http.StatusNotFound,
		},
		{
			name:           "store_error",
			storeErr:       errors.New("error"),
			postID:         "1",
			expectedStatus: http.StatusInternalServerError,
		},
		{
			name:           "invalid_id",
			storeErr:       nil,
			postID:         "abc",
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := &MockStore{
				RemovePostFunc: func(pid int) error {
					if pid != 1 {
						return store.ErrNotFound
					}

					return tt.storeErr
				},
			}

			server := New(store, &MockAIClient{})
			req := httptest.NewRequest(
				http.MethodDelete,
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
