package cache

import (
	"context"
	"crm_system/internal/crm_core/entity"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"time"
)

type Contact interface {
	Get(ctx context.Context, key string) (*entity.Contact, error)
	Set(ctx context.Context, key string, value *entity.Contact) error
}

type ContactCache struct {
	Expiration time.Duration
	redisCli   *redis.Client
}

func NewContactCache(redisCli *redis.Client, expiration time.Duration) Contact {
	return &ContactCache{
		redisCli:   redisCli,
		Expiration: expiration,
	}
}

func (c *ContactCache) Get(ctx context.Context, key string) (*entity.Contact, error) {
	value := c.redisCli.Get(ctx, key).Val()

	if value == "" {
		return nil, nil
	}

	var contact *entity.Contact

	err := json.Unmarshal([]byte(value), &contact)
	if err != nil {
		return nil, err
	}

	return contact, nil
}

func (c *ContactCache) Set(ctx context.Context, key string, value *entity.Contact) error {
	contactJson, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return c.redisCli.Set(ctx, key, string(contactJson), c.Expiration).Err()
}
