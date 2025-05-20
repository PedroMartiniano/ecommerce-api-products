package repositories

import (
	"context"
	"time"
)

type IRedisRepository interface {
	// value must be a pointer to the type of the value to be returned
	Get(ctx context.Context, key string, value any) error

	Set(ctx context.Context, key string, value any, expiration time.Duration) error

	Delete(ctx context.Context, key string) error
}
