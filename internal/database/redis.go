package database

import (
	"context"
	"fmt"
	"log"

	"server/pkg/config"

	"github.com/redis/go-redis/v9"
)

var Ctx = context.Background()

func NewRedis(cfg *config.Config) *redis.Client {

	r := cfg.Redis

	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", r.Host, r.Port),
	})

	log.Println("redis connected")

	return client
}
