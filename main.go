package main

import (
	"fmt"
	"go-board/mongodb"
	"go-board/router"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env")
	}
	mongodb.Init(
		os.Getenv("MONGO_URL"),
		os.Getenv("MONGO_DB"),
		os.Getenv("MONGO_COLLECTION"))
	router.Init()
	fmt.Println("http://localhost:8080")
	http.ListenAndServe(":8080", router.R)
}
