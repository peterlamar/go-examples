package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	query := `select * from cs_ocean_contract_global_gri
where $1::date between from_date and to_date`
// Get data in cache if exists
// else get from db

// match env id, carrier id, trade lane, container type

	err := client.Set(query, "value1", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get(query).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := client.Get("key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist

}
