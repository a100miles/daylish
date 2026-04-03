package models

type Post struct {
	ID          int    `json:"id"`
	ImageID     string `json:"imageID"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

var Posts = make(map[int]Post)

var IDCounter = 1
