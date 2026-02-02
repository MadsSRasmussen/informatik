package store

import (
	"informatik/api/internal/db"
)

type Storer interface {
	ListPosts() ([]db.Post, error)
	GetPost(pid int) (*db.Post, error)
	CreatePost(content string) (*PostCreateResponse, error)
	RemovePost(pid int) error
}
