package handlers

import(
	"github.com/redis/go=redis/v9"
	"go_redisCRUD/database"
)

func RedisTypes(rdb *redis.Client){
	//String

	rdb.Set(database.Ctx, "name", "Alice",0)
	StringVal,_:= rdb.Get(database.Ctx, "name").Result()
	
	fmt.Printf("The String mapped Value is %s\n", StringVal)


	//Hash

	rdb.HSet(database.Ctx, "user:1", "name", "Alice", "age", "30")
	hashVal, _:= rdb.HGetAll(database.Ctx, "user:1").Result()

	fmt.Printf("The HashMap values %s\n", hashVal)



}
