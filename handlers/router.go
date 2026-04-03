package handlers

import "github.com/gorilla/mux"

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api/posts", CreatePost).Methods("POST")
	r.HandleFunc("/api/posts", GetPosts).Methods("GET")
	r.HandleFunc("/api/posts/{id}", GetPost).Methods("GET")
	r.HandleFunc("/api/posts/{id}", UpdatePost).Methods("PUT")
	r.HandleFunc("/api/posts/{id}", DeletePost).Methods("DELETE")

	return r
}