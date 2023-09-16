package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
)

var _ Interface = (*Cache)(nil)

func New(rdb *redis.Client) *Cache {
	return &Cache{
		rdb: rdb,
	}
}

type Cache struct {
	rdb *redis.Client
}

func (c *Cache) Create(ctx context.Context, key string, v any) error {
	var (
		val []byte
		err error
	)

	if val, err = json.Marshal(v); err != nil {
		return fmt.Errorf("create / marshall :%w", err)
	}

	cmd := c.rdb.Set(ctx, key, val, 0)
	if cmd.Err() != nil {
		return fmt.Errorf("create / set :%w", cmd.Err())
	}

	log.Println("key", key, "cache record created")

	return nil
}

func (c *Cache) Get(ctx context.Context, key string, v any) error {
	cmd := c.rdb.Get(ctx, key)

	if cmd.Err() != nil {
		fmt.Println(cmd.Err().Error())
		return fmt.Errorf("get / get :%w", cmd.Err())
	}

	var (
		val []byte
		err error
	)

	if val, err = cmd.Bytes(); err != nil {
		return fmt.Errorf("get / bytes : %w", err)

	}

	if err = json.Unmarshal(val, v); err != nil {
		return fmt.Errorf("get / unmarshal : %w", err)
	}

	log.Println("key", key, "cache record fetched")

	return nil
}

func (c *Cache) Delete(ctx context.Context, key string) error {
	if err := c.rdb.Del(ctx, key).Err(); err != nil {
		return fmt.Errorf("delete / del : %w", err)
	}

	log.Println("key", key, "cache record deleted")
	return nil
}

func (c *Cache) List(ctx context.Context, key string, v any) error {
	cmd := c.rdb.Get(ctx, key)
	if cmd.Err() != nil {
		return fmt.Errorf("list /get :%w", cmd.Err())
	}

	val, err := cmd.Bytes()
	if err != nil {
		return fmt.Errorf("list /bytes :%w", err)
	}

	if err := json.Unmarshal(val, v); err != nil {
		return fmt.Errorf("list / unmarshal :%w", err)
	}

	log.Println(key, "cache list fetch")

	return nil
}
