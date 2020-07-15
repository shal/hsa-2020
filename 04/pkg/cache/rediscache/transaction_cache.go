package rediscache

import (
	"context"
	"encoding/json"
	"github.com/shal/hsa-2020/04/pkg/cache"

	"github.com/gomodule/redigo/redis"

	"github.com/shal/hsa-2020/04/pkg/model"
)

type TransactionCache struct {
	cache *Cache
}

func NewTransactionRepository(cache *Cache) *TransactionCache {
	return &TransactionCache{
		cache: cache,
	}
}

func (r *TransactionCache) Set(ctx context.Context, result *model.Result) error {
	conn, err := r.cache.pool.DialContext(ctx)
	if err != nil {
		return err
	}

	data, err := json.Marshal(result)
	if err != nil {
		return err
	}

	_, err = conn.Do("SET", "result", data, "EX", 30)
	if err != nil {
		return err
	}

	return nil
}

func (r *TransactionCache) Get(ctx context.Context) (*model.Result, error) {
	conn, err := r.cache.pool.DialContext(ctx)
	if err != nil {
		return nil, err
	}

	data, err := redis.Bytes(conn.Do("GET", "result"))
	if err == redis.ErrNil {
		return nil, cache.ErrNotFound
	}

	if err != nil {
		return nil, err
	}

	_ = conn.Close()

	payload := model.Result{}
	if err := json.Unmarshal(data, &payload); err != nil {
		return nil, err
	}

	return &payload, nil
}
