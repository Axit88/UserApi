package awsClient

import (
	"context"
	"time"

	"github.com/Axit88/UserApi/src/constants"
	"github.com/Axit88/UserApi/src/domain/userService/core/ports/outgoing"
	"github.com/MindTickle/mt-go-logger/logger"
	"github.com/go-redis/redis/v8"
)

type RedisImpl struct {
	logger *logger.LoggerImpl
	redis  *redis.Client
}

func NewRedisClient(l *logger.LoggerImpl) outgoing.RedisClient {

	if constants.IsMock {
		return RedisMockClient{}
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr: "redisEndpoint", // Redis server address
		DB:   0,               // Redis database number
	})

	return &RedisImpl{
		logger: l,
		redis:  redisClient,
	}
}

func (client RedisImpl) RedisSetkey(key string, value string, expiryTime time.Duration) error {
	err := client.redis.Set(context.Background(), key, value, expiryTime).Err()
	return err
}
