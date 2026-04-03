package handlers

import (
	"daylish/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	var post models.Post
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	post.ID = models.IDCounter
	models.IDCounter++
	models.Posts[post.ID] = post
	json.NewEncoder(w).Encode(post)
}

func GetPosts(w http.ResponseWriter, r *http.Request) {
	var postList []models.Post
	for _, post := range models.Posts {
		postList = append(postList, post)
	}
	json.NewEncoder(w).Encode(postList)
}

func GetPost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil || models.Posts[id].ID == 0 {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	} 
	json.NewEncoder(w).Encode(models.Posts[id])
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil || models.Posts[id].ID == 0 {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}
	var updatedPost models.Post
	if err := json.NewDecoder(r.Body).Decode(&updatedPost); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedPost.ID = id
	models.Posts[id] = updatedPost
	json.NewEncoder(w).Encode(updatedPost)
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil || models.Posts[id].ID == 0 {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}
	delete(models.Posts, id)
	w.WriteHeader(http.StatusNoContent)
}
