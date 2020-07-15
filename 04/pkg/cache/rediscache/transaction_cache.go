package rediscache

import (
	"context"
	"encoding/json"

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

func (r *TransactionCache) Set(ctx context.Context, transaction *model.Transaction) error {
	conn, err := r.cache.pool.DialContext(ctx)
	if err != nil {
		return err
	}

	data, err := json.Marshal(transaction)
	if err != nil {
		return err
	}

	_, err = conn.Do("SET", transaction.ID, data, "EX", 10)
	if err != nil {
		return err
	}

	return nil
}

func (r *TransactionCache) Get(ctx context.Context, id string) (*model.Transaction, error) {
	conn, err := r.cache.pool.DialContext(ctx)
	if err != nil {
		return nil, err
	}

	data, err := redis.Bytes(conn.Do("GET", id))
	if err == redis.ErrNil {
		return nil, cache.ErrNotFound
	}

	if err != nil {
		return nil, err
	}

	_ = conn.Close()

	payload := model.Transaction{}
	if err := json.Unmarshal(data, &payload); err != nil {
		return nil, err
	}

	return &payload, nil
}
