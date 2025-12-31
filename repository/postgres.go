package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/hexa-SaikumarAilwar/RedisPOC.git/entity"

	_ "github.com/lib/pq" // PostgreSQL driver
)

type repo struct {
	db *sql.DB
}

func NewPostRepository(connStr string) (PostRepository, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	// Verify connection
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	log.Println("Successfully connected to PostgreSQL")

	return &repo{db: db}, nil
}

func (r *repo) Save(post *entity.Post) (*entity.Post, error) {
	ctx := context.Background()
	query := `INSERT INTO posts (title, text) VALUES ($1, $2) RETURNING id`
	err := r.db.QueryRowContext(ctx, query, post.Title, post.Text).Scan(&post.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to save post: %v", err)
	}

	return post, nil
}

func (r *repo) FindAll() ([]entity.Post, error) {
	ctx := context.Background()

	query := `SELECT id, title, text FROM posts`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch posts: %v", err)
	}
	defer rows.Close()

	var posts []entity.Post
	for rows.Next() {
		var post entity.Post
		if err := rows.Scan(&post.Id, &post.Title, &post.Text); err != nil {
			return nil, fmt.Errorf("failed to scan post: %v", err)
		}
		posts = append(posts, post)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %v", err)
	}

	return posts, nil
}

func (r *repo) FindById(id int) (*entity.Post, error) {
	ctx := context.Background()
	query := `SELECT id, title, text FROM posts WHERE id = $1`

	var post entity.Post
	err := r.db.QueryRowContext(ctx, query, id).Scan(&post.Id, &post.Title, &post.Text)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("post with id %d not found", id)
		}
		return nil, fmt.Errorf("failed to fetch post: %v", err)
	}

	return &post, nil
}
