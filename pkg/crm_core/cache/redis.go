package cache

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

func NewRedisClient() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	//err := client.Ping(context.Background()).Err()
	//if err != nil {
	//	return nil, err
	//}

	fmt.Println(client.Ping(context.Background()).Err())

	return client, nil
}
