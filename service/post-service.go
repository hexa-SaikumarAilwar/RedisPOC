package service

import (
	"github.com/hexa-SaikumarAilwar/RedisPOC.git/POC/entity"
	"github.com/hexa-SaikumarAilwar/RedisPOC.git/POC/repository"
	"errors"
)

type PostService interface {
	Validate(post *entity.Post) error
	CreatePost(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}

type service struct{}

var (
	repo repository.PostRepository
)

func NewPostService(r repository.PostRepository) PostService {
	repo = r
	return &service{}
}

func (*service) Validate(post *entity.Post) error {
	if post == nil {
		err := errors.New("the Post is empty")
		return err
	}
	if post.Title == "" {
		err := errors.New("the post title is empty")
		return err
	}
	return nil
}
func (*service) CreatePost(post *entity.Post) (*entity.Post, error) {
	return repo.Save(post)
}
func (*service) FindAll() ([]entity.Post, error) {
	return repo.FindAll()
}
