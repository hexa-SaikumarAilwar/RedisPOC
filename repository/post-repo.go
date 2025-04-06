package repository

import "github.com/hexa-SaikumarAilwar/RedisPOC.git/entity"

type PostRepository interface {
	Save(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
	FindById(id int)(*entity.Post, error)
}
