package handlers

import(
	"github.com/redis/go-redis/v9"
	"go_redisCRUD/database"
	"fmt"
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

	//List
	rdb.RPush(database.Ctx, "tasks", "task1", "task2")
	listVal, _:= rdb.LPop(database.Ctx, "tasks").Result()

	fmt.Printf("The list Values are %s\n", listVal)

	//Set
	rdb.SAdd(database.Ctx, "skills", "go", "Java")
	setVal, _:= rdb.SMembers(database.Ctx, "skills").Result()

	fmt.Printf("THe Set values are %s\n", setVal)

	//SortedSet

	rdb.ZAdd(database.Ctx, "leaderboard", redis.Z{Score: 100,Member: "Alice"})
	rdb.ZAdd(database.Ctx, "leaderboard", redis.Z{Score: 120,Member:"Bob"})

	leader, _:= rdb.ZRangeWithScores(database.Ctx, "leaderboard", 0, -1).Result()
	fmt.Println("Leaderboard:")
	for _, z:= range leader{
		fmt.Printf("The value of ZAdd is %v : %v\n", z.Score, z.Member)
	}




}
