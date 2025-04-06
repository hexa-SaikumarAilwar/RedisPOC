package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/hexa-SaikumarAilwar/RedisPOC.git/cache"
	"github.com/hexa-SaikumarAilwar/RedisPOC.git/entity"
	"github.com/hexa-SaikumarAilwar/RedisPOC.git/service"
)

var (
	postService service.PostService
	postCache   cache.PostCache
)

type PostController interface {
	GetPosts(resp http.ResponseWriter, req *http.Request)
	GetPostById(resp http.ResponseWriter, req *http.Request)
	AddPost(resp http.ResponseWriter, req *http.Request)
}

type controller struct{}

func NewPostController(service service.PostService, cache cache.PostCache) PostController {
	postService = service
	postCache = cache
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

func (*controller) GetPostById(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")

	// Extract ID from URL params
	vars := mux.Vars(req)
	id := vars["id"]

	// Convert to int (with error handling)
	postId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(resp, `{"error": "Invalid post ID"}`, http.StatusBadRequest)
		return
	}

	var post *entity.Post = postCache.Get(id)
	if post == nil {
		post, err := postService.FindById(postId)
		if err != nil {
			http.Error(resp, `{"error": "Post not found"}`, http.StatusNotFound)
			return
		}
		postCache.Set(id, post)
		json.NewEncoder(resp).Encode(post)
	} else {
		json.NewEncoder(resp).Encode(post)
	}
}
