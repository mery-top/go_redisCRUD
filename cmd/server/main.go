package main

import(
	"go_redisCRUD/database"
	"go_redisCRUD/handlers"
)

func main(){
	rdb:= database.ConnectRedis()
	handlers.RedisCRUD(rdb)
}
