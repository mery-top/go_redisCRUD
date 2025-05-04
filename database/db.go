package database


import(
	"log"
	"context"
	"github.com/redis/go-redis/v9"
)


var ctx = context.Background()

func ConnectRedis() *redis.Client{
	rdb:= redis.NewClient(&redis.Options{
		Addr: "redis:6379",
		DB: 0,
	})

	return rdb
}
