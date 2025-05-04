package handlers

import(
	"go_redisCRUD/database"
	"github.com/redis/go-redis/v9"
	"fmt"
	"time"
)

func RedisCRUD(rdb *redis.Client){
	sessionID:= "user1234"
	userID:= "101"

	err:= rdb.Set(database.Ctx, sessionID, userID, 10*time.Minute).Err()
	if err!=nil{
		panic(err)
	}

	fmt.Println("Session set successfully!")

	val, err:= rdb.Get(database.Ctx, sessionID).Result()
	if err!= nil{
		panic(err)
	}

	fmt.Printf("The session value is %s\n", val)

	rdb.Del(database.Ctx, sessionID)
	fmt.Println("Session deleted successfully")
}
