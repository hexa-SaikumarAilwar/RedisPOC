package main

import (
	"log"

	"github.com/hexa-SaikumarAilwar/RedisPOC.git/cache"
	"github.com/hexa-SaikumarAilwar/RedisPOC.git/controller"
	router "github.com/hexa-SaikumarAilwar/RedisPOC.git/http"
	"github.com/hexa-SaikumarAilwar/RedisPOC.git/repository"
	"github.com/hexa-SaikumarAilwar/RedisPOC.git/service"
)

const (
	connStr = "postgres://postgres:saikumar@localhost:5432/postgres?sslmode=disable"
	port    = ":4000"
)

func main() {
	// Intialize router
	var httpRouter router.Router = router.NewMuxRouter()
	// var httpRouter router.Router = router.NewGinRouter()
	var postCache cache.PostCache = cache.NewRedisCache("localhost:6379", 0, 10)
	// var postCache cache.PostCache = cache.NewValkeyCache("localhost:6379", 10*time.Second)

	// Initialize repository
	repo, err := repository.NewPostRepository(connStr)
	if err != nil {
		log.Fatalf("Failed to initialize repository: %v", err)
	}

	// Initialize service and controller
	postService := service.NewPostService(repo, postCache)
	postController := controller.NewPostController(postService)

	// API routes
	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.GET("/post/{id}", postController.GetPostById)
	httpRouter.POST("/posts", postController.AddPost)

	httpRouter.SERVE(port)
}
