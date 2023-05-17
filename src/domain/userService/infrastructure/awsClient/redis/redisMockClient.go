package awsClient

import (
	"time"
)

type RedisMockClient struct {
}

func (client RedisMockClient) RedisSetkey(key string, value string, expiryTime time.Duration) error {
	return nil
}
