package cache

import (
	"context"

	"github.com/shal/hsa-2020/04/pkg/model"
)

type TransactionCache interface {
	Set(ctx context.Context, transaction *model.Transaction) error
	Get(ctx context.Context, id string) (*model.Transaction, error)
}
