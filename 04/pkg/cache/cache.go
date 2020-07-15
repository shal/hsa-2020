package cache

type Cache interface {
	Transaction() TransactionCache
}
