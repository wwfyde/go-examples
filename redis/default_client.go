package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

func main() {

	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})
	//// Read option with URL
	// opt, _ := redis.ParseURL("")
	//ctx := context.Background()
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Execute commands
	val := rdb.Get(ctx, "a")
	fmt.Println(val.Val(), val.Err())

	// Execute unsupported commands
	val2, _ := rdb.Do(ctx, "get", "a").Result()
	fmt.Println(val2.(string))
	// Execute commands and Check if key not exist
	val3, err := rdb.Get(ctx, "aa").Result()
	{
		switch {
		case err == redis.Nil:
			fmt.Println("key does not exist")
		case err != nil:
			fmt.Println("Get failed: ", err)
		case val3 == "":
			fmt.Println("value is empty")
		default:
			fmt.Println("get value: ", val3)
		}
	}

}
