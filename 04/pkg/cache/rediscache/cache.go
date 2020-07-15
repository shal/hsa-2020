package rediscache

import (
	"context"
	"time"

	"github.com/gomodule/redigo/redis"

	"github.com/shal/hsa-2020/04/pkg/cache"
	"github.com/shal/hsa-2020/04/pkg/config"
)

type Cache struct {
	pool *redis.Pool

	transactionCache *TransactionCache
}

func New(ctx context.Context, conf config.Cache) (*Cache, error) {
	pool := redis.Pool{
		MaxIdle:   50,
		MaxActive: 10000,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", conf.Addr(),
				redis.DialPassword(conf.Password),
				redis.DialDatabase(conf.Database),
				redis.DialConnectTimeout(10*time.Second),
			)

			// Connection error handling.
			if err != nil {
				return nil, err
			}
			return conn, err
		},
	}

	conn, err := pool.DialContext(ctx)
	if err != nil {
		return nil, err
	}

	if err := conn.Close(); err != nil {
		return nil, err
	}

	return &Cache{
		pool: &pool,
	}, nil
}

func (c *Cache) Transaction() cache.TransactionCache {
	if c.transactionCache == nil {
		c.transactionCache = NewTransactionRepository(c)
	}

	return c.transactionCache
}
