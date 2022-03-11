package main

import (
	"context"
	redis "github.com/go-redis/redis/v8"
	"go-board/router"
	"net/http"
)

var ctx = context.Background()

func redisTest() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "49.50.162.177:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := rdb.Set(ctx, "kkey", "vvalue", 0).Err()
	if err != nil {
		panic(err)
	}
}

func main() {
	router.Init()
	redisTest()
	http.ListenAndServe(":8080", router.R)
}
