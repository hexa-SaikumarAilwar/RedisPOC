package controller

import (
	"encoding/json"
	"net/http"

	"github.com/hexa-SaikumarAilwar/RedisPOC.git/entity"
	"github.com/hexa-SaikumarAilwar/RedisPOC.git/service"
)

type PostController interface {
	GetPosts(resp http.ResponseWriter, req *http.Request)
	GetPostById(resp http.ResponseWriter, req *http.Request)
	AddPost(resp http.ResponseWriter, req *http.Request)
}

type controller struct {
	postService service.PostService
}

func NewPostController(s service.PostService) PostController {
	return &controller{postService: s}
}

func (c *controller) GetPosts(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")

	posts, err := c.postService.FindAll()
	if err != nil {
		http.Error(resp, `{"error": "Failed to fetch posts"}`, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(resp).Encode(posts)
}

func (c *controller) AddPost(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")

	var post entity.Post
	if err := json.NewDecoder(req.Body).Decode(&post); err != nil {
		http.Error(resp, `{"error": "Failed to decode post"}`, http.StatusInternalServerError)
		return
	}

	err := c.postService.Validate(&post)
	if err != nil {
		http.Error(resp, `{"error": "Failed to validate post"}`, http.StatusInternalServerError)
		return
	}
	savedPost, err := c.postService.CreatePost(&post)
	if err != nil {
		http.Error(resp, `{"error": "Failed to save post"}`, http.StatusInternalServerError)
		return
	}

	resp.WriteHeader(http.StatusCreated)
	json.NewEncoder(resp).Encode(savedPost)
}

func (c *controller) GetPostById(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")

	postId := req.Context().Value("id").(int)

	post, err := c.postService.FindById(postId)
	if err != nil {
		http.Error(resp, `{"error": "Post not found"}`, http.StatusNotFound)
		return
	}

	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(post)
}
