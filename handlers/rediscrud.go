package handlers

import(
	"go_redisCRUD/database"
	"fmt"
	"context"
	"time"
)

func redisCRUD(rdb *redis.Client){
	sessionID:= "user1234"
	userID:= "101"

	err:= rdb.Set(database.ctx, sessionID, userID, 10*time.Minute).Err()
	if err!=nil{
		panic(err)
	}

	fmt.Println("Session set successfully!")

	val, err:= rdb.Get(database.ctx, sessionID).Result()
	if err!= nil{
		panic(er)
	}

	fmt.Printf("The session value is %s", val)

	rdb.Del(database.ctx, sessionID)

	fmt.Println("Session deleted successfully")
}
