package cache

import "context"

type Interface interface {
	Create(ctx context.Context, key string, v any) error
	Get(ctx context.Context, key string, v any) error
	Delete(ctx context.Context, key string) error
	List(ctx context.Context, key string, v any) error
}
