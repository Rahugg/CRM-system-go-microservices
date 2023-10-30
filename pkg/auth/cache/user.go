package cache

import (
	"context"
	"crm_system/internal/auth/entity"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"time"
)

type User interface {
	Get(ctx context.Context, key string) (*entity.User, error)
	Set(ctx context.Context, key string, value *entity.User) error
}

type UserCache struct {
	Expiration time.Duration
	redisCli   *redis.Client
}

func NewUserCache(redisCli *redis.Client) User {
	return &UserCache{
		redisCli: redisCli,
	}
}

func (u *UserCache) Get(ctx context.Context, key string) (*entity.User, error) {
	value := u.redisCli.Get(ctx, key).Val()

	var user *entity.User

	err := json.Unmarshal([]byte(value), &user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserCache) Set(ctx context.Context, key string, value *entity.User) error {
	userJson, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return u.redisCli.Set(ctx, key, string(userJson), u.Expiration).Err()
}
