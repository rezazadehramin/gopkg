package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type rdis struct {
	Client *redis.Client
	Config Configurator
}

// Get retrieves a value from the cache storage and it also
// provides an error if exists.
func (r *rdis) Get(k string) (interface{}, error) {
	return r.Client.Get(context.Background(), k).Result()
}

// Get retrieves a value from the cache storage if an error
// is retrieved then an empty interface will be returned.
func (r *rdis) GetSafe(k string) interface{} {
	res, err := r.Client.Get(context.Background(), k).Result()
	if err == redis.Nil {
		return ""
	} else if err != nil {
		return nil
	}

	return res

}

// Set puts a value on the cache storage, ttls must be an amount
// of seconds > 0.
//
// If ttl is equal to 0 then the "forever" persistance policy is
// used when supported by the engine.
func (r *rdis) Set(k string, v interface{}, ttl time.Duration) error {
	return r.Client.Set(context.Background(), k, v, ttl).Err()
}

// NewRedis constructs a redis cache client.
func NewRedis(config Configurator, client *redis.Client) Engine {
	if config == nil {
		config = NewENVConfig()
	}

	if client == nil {
		client = rootRedisClient(config)
	}

	return &rdis{
		Client: client,
		Config: config,
	}
}

func rootRedisClient(config Configurator) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     config.Addr(),
		Password: config.Pwd(),
	})
}
