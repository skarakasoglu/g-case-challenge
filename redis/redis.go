// Package redis creates new redis client by using the connection url.
package redis

import (
	"github.com/go-redis/redis/v8"
	"log"
)

func NewClient(url string) *redis.Client {
	opts, err := redis.ParseURL(url)
	if err != nil {
		log.Fatalf("error on parsing redis URL: %v", err)
	}

	rdb := redis.NewClient(opts)
	return rdb
}
