package cache

import "errors"

var (
	ErrNotFound = errors.New("not found")
)

type Cache interface {
	Transaction() TransactionCache
}
