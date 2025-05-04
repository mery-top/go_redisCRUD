package database


import(
	"context"
	"github.com/redis/go-redis/v9"
)


var Ctx = context.Background()

func ConnectRedis() *redis.Client{
	rdb:= redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB: 0,
	})

	return rdb
}
