package main

import (
	"daylish/handlers"
	"fmt"
	"net/http"
)

func main() {
	r := handlers.NewRouter()
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Println(err)
	}

}

//need a db that will store posts post.go
//authorization & authenthication system.
//
