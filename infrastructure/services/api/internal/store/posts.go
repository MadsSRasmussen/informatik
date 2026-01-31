package store

import (
	"database/sql"
	"errors"
	"informatik/api/internal/db"
)

func (s *Store) ListPosts() ([]db.Post, error) {
	rows, err := s.db.Query("SELECT * FROM posts")
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return []db.Post{}, nil
		}
		return nil, err
	}
	defer rows.Close()

	var posts []db.Post

	for rows.Next() {
		var post db.Post

		if err := rows.Scan(
			&post.ID,
			&post.Content,
			&post.CreatedAt,
		); err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

func (s *Store) GetPost(pid int) (*db.Post, error) {
	row := s.db.QueryRow(
		"SELECT * FROM posts WHERE id = ?",
		pid,
	)

	var post db.Post
	if err := row.Scan(
		&post.ID,
		&post.Content,
		&post.CreatedAt,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return &post, nil
}

type PostCreateResponse struct {
	ID int64 `json:"id"`
}

func (s *Store) CreatePost(content string) (*PostCreateResponse, error) {
	res, err := s.db.Exec(
		"INSERT INTO posts (content) VALUES (?)",
		content,
	)
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &PostCreateResponse{
		ID: id,
	}, nil
}
