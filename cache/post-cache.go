package cache

import "github.com/hexa-SaikumarAilwar/RedisPOC.git/entity"

type PostCache interface {
	Set(key string, value *entity.Post)
	Get(key string) *entity.Post
}