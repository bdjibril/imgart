package redis

import (
	"time"

	"github.com/go-redis/redis"
	"github.com/talento90/gorpo/pkg/cache"
	"github.com/talento90/gorpo/pkg/errors"
)

func handleError(err error) error {
	if err == nil {
		return nil
	}

	if err == redis.Nil {
		return errors.ENotExists("Item does not exists", err)
	}

	return errors.EInternal("Error occured", err)
}

type redisRepository struct {
	client *redis.Client
}

func NewRedisRepository(client *redis.Client) cache.Cache {
	return &redisRepository{client: client}
}

func (r *redisRepository) Get(key string) ([]byte, error) {
	result, err := r.client.Get(key).Bytes()

	return result, handleError(err)
}

func (r *redisRepository) Set(key string, value []byte, expiration time.Duration) error {
	err := r.client.Set(key, value, expiration).Err()

	return handleError(err)
}