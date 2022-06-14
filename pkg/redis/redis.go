package redis

import (
	"github.com/sm0kernone/crm-common/pkg/config"
	"fmt"
	"github.com/go-redis/redis"
	"strconv"
)

// New instantiates new Redis client
func New(cfg config.Redis) (*redis.Client, error) {
	opts := redis.Options{
		Addr:     cfg.Address + ":" + strconv.Itoa(cfg.Port),
		DB:       0,
		Password: cfg.Password,
	}

	client := redis.NewClient(&opts)

	_, err := client.Ping().Result()
	if err != nil {
		return nil, fmt.Errorf("cannot connect to Redis Addr %v, Port %v Reason %v", cfg.Address, cfg.Port, err)
	}

	return client, nil
}
