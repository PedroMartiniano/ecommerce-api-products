package repositories

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	pr "github.com/PedroMartiniano/ecommerce-api-products/internal/application/ports/repositories"
	"github.com/PedroMartiniano/ecommerce-api-products/internal/configs"
	"github.com/redis/go-redis/v9"
)

var logger = configs.GetLogger()

type redisRepository struct {
	redis *redis.Client
}

func NewRedisRepository(redis *redis.Client) pr.IRedisRepository {
	return &redisRepository{
		redis: redis,
	}
}

func (r *redisRepository) Get(ctx context.Context, key string, value any) error {
	data, err := r.redis.Get(ctx, key).Bytes()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return configs.NewError(configs.ErrNotFound, err)
		}
		return configs.NewError(configs.ErrInternalServer, err)
	}

	err = json.Unmarshal(data, value)
	if err != nil {
		return configs.NewError(configs.ErrInternalServer, fmt.Errorf("failed to unmarshal data: %w", err))
	}

	return nil
}

func (r *redisRepository) Set(ctx context.Context, key string, value any, expiration time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		return configs.NewError(configs.ErrInternalServer, fmt.Errorf("failed to marshal value: %w", err))
	}

	err = r.redis.Set(ctx, key, data, expiration).Err()
	if err != nil {
		return configs.NewError(configs.ErrInternalServer, err)
	}

	return nil
}

func (r *redisRepository) Delete(ctx context.Context, key string) error {
	err := r.redis.Del(ctx, key).Err()
	if err != nil {
		return configs.NewError(configs.ErrInternalServer, err)
	}

	return nil
}
