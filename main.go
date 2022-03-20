package main

import (
	"fmt"
	"go-board/mongodb"
	"go-board/router"
	"net/http"
)

func main() {
	mongodb.Init(
		"mongodb://admin:jinho9611%40@118.67.143.181:27017/admin",
		"board",
		"articles")
	router.Init()
	fmt.Println("http://localhost:8080")
	http.ListenAndServe(":8080", router.R)
}
