package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

func main() {

	redisdb := redis.NewClient(&redis.Options{
		Addr: ":6379",
	})
	err := redisdb.Set("name", "lex", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := redisdb.Get("name").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("name", val)

	val2, err := redisdb.Get("missing_key").Result()
	if err == redis.Nil {
		fmt.Println("missing_key does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("missing_key", val2)
	}
}
