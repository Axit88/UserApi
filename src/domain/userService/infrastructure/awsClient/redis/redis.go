package awsClient

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func RedisSet(client *redis.Client, redisEndpoint string) error {

	err := client.Set(context.Background(), "name", "jay", 0).Err()
	if err != nil {
		return err
	}

	val, err := client.Get(context.Background(), "name").Result()
	if err != nil {
		return err
	}
	fmt.Println("key", val)
	return nil
}