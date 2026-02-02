package server

import (
	"informatik/api/internal/ai"
	"informatik/api/internal/db"
	"informatik/api/internal/store"
)

type MockStore struct {
	ListPostsFunc  func() ([]db.Post, error)
	GetPostFunc    func(pid int) (*db.Post, error)
	CreatePostFunc func(content string) (*store.PostCreateResponse, error)
	RemovePostFunc func(pid int) error
}

func (m *MockStore) ListPosts() ([]db.Post, error) {
	return m.ListPostsFunc()
}

func (m *MockStore) GetPost(pid int) (*db.Post, error) {
	return m.GetPostFunc(pid)
}

func (m *MockStore) CreatePost(content string) (*store.PostCreateResponse, error) {
	return m.CreatePostFunc(content)
}

func (m *MockStore) RemovePost(pid int) error {
	return m.RemovePostFunc(pid)
}

type MockAIClient struct {
	GetCompletionFunc func(r *ai.CompletionsRequestBody) (*ai.Message, error)
}

func (m *MockAIClient) GetCompletion(r *ai.CompletionsRequestBody) (*ai.Message, error) {
	return m.GetCompletionFunc(r)
}
