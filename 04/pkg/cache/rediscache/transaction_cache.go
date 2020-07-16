package rediscache

import (
	"context"
	"encoding/json"
	"log"
	"math/rand"
	"time"

	"github.com/gomodule/redigo/redis"

	"github.com/shal/hsa-2020/04/pkg/cache"
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

	err = conn.Send("SET", "result", data, "EX", 30)
	if err != nil {
		return err
	}

	resultTime := time.Now().Add(30 * time.Second).Format(time.RFC3339)
	err = conn.Send("SET", "result_time", resultTime, "EX", 40)
	if err != nil {
		return err
	}

	if err := conn.Flush(); err != nil {
		return err
	}

	if _, err := conn.Receive(); err != nil {
		return err
	}

	return nil
}

func (r *TransactionCache) Get(ctx context.Context) (*model.Result, error) {
	conn, err := r.cache.pool.DialContext(ctx)
	if err != nil {
		return nil, err
	}

	timeData, err := redis.String(conn.Do("GET", "result_time"))
	if err == redis.ErrNil {
		return nil, cache.ErrNotFound
	}

	resultTime, err := time.Parse(time.RFC3339, timeData)
	if err == redis.ErrNil {
		return nil, cache.ErrNotFound
	}

	if err != nil {
		return nil, err
	}

	// Probabilistic cache flushing.
	left := resultTime.Sub(time.Now())
	if left < 0 {
		return nil, cache.ErrNotFound
	} else if left < time.Second {
		if rand.Float64() <= 1.0-left.Seconds() {

			return nil, cache.ErrNotFound
		}
		log.Println("not dispatched update")
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
