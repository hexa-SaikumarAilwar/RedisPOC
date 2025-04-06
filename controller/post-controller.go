package controller

import (
	"github.com/hexa-SaikumarAilwar/RedisPOC.git/POC/entity"
	"github.com/hexa-SaikumarAilwar/RedisPOC.git/POC/service"
	"encoding/json"
	"net/http"
)

var (
	postService service.PostService
)

type PostController interface {
	GetPosts(resp http.ResponseWriter, req *http.Request)
	AddPost(resp http.ResponseWriter, req *http.Request)
}

type controller struct{}

func NewPostController(service service.PostService) PostController {
	postService = service
	return &controller{}
}

func (*controller) GetPosts(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")

	posts, err := postService.FindAll()
	if err != nil {
		http.Error(resp, `{"error": "Failed to fetch posts"}`, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(resp).Encode(posts)
}

func (*controller) AddPost(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")

	var post entity.Post
	if err := json.NewDecoder(req.Body).Decode(&post); err != nil {
		http.Error(resp, `{"error": "Failed to decode post"}`, http.StatusInternalServerError)
		return
	}

	err := postService.Validate(&post)
	if err != nil {
		http.Error(resp, `{"error": "Failed to validate post"}`, http.StatusInternalServerError)
		return
	}
	savedPost, err := postService.CreatePost(&post)
	if err != nil {
		http.Error(resp, `{"error": "Failed to save post"}`, http.StatusInternalServerError)
		return
	}

	resp.WriteHeader(http.StatusCreated)
	json.NewEncoder(resp).Encode(savedPost)
}
