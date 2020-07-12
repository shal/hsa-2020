package store

import (
	"context"

	"github.com/shal/hsa-2020/04/pkg/model"
)

type TransactionRepository interface {
	Create(ctx context.Context, transaction *model.Transaction) error
	FindByID(ctx context.Context, id string) (*model.Transaction, error)
	All(ctx context.Context) ([]model.Transaction, error)
}
