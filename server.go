package main

import (
	"log"

	"github.com/hexa-SaikumarAilwar/RedisPOC.git/controller"
	router "github.com/hexa-SaikumarAilwar/RedisPOC.git/http"
	"github.com/hexa-SaikumarAilwar/RedisPOC.git/repository"
	"github.com/hexa-SaikumarAilwar/RedisPOC.git/service"
)

const (
	connStr = "postgres://postgres:saikumar@localhost:5432/postgres?sslmode=disable"
	port    = ":4000"
)

var (
	httpRouter router.Router = router.NewMuxRouter()
)

func main() {
	// Initialize repository
	repo, err := repository.NewRepository(connStr)
	if err != nil {
		log.Fatalf("Failed to initialize repository: %v", err)
	}

	// Initialize service and controller
	postService := service.NewPostService(repo)
	postController := controller.NewPostController(postService)

	// API routes
	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.POST("/posts", postController.AddPost)

	httpRouter.SERVE(port)
}
