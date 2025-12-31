package service

import (
	"errors"
	"strconv"

	"github.com/hexa-SaikumarAilwar/RedisPOC.git/cache"
	"github.com/hexa-SaikumarAilwar/RedisPOC.git/entity"
	"github.com/hexa-SaikumarAilwar/RedisPOC.git/repository"
)

type PostService interface {
	Validate(post *entity.Post) error
	CreatePost(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
	FindById(id int) (*entity.Post, error)
}

type service struct{
	repo repository.PostRepository
	cache cache.PostCache
}

func NewPostService(r repository.PostRepository, c cache.PostCache) PostService {
	return &service{repo: r, cache: c}
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
func (s *service) CreatePost(post *entity.Post) (*entity.Post, error) {
	return s.repo.Save(post)
}
func (s *service) FindAll() ([]entity.Post, error) {
	return s.repo.FindAll()
}

func (s *service) FindById(id int) (*entity.Post, error) {
	key := strconv.Itoa(id)

	post := s.cache.Get(key)
	if post != nil {
		return post, nil
	}

	post, err := s.repo.FindById(id)
	if err != nil {
		return nil, err
	}

	s.cache.Set(key, post)
	return post, nil
}
