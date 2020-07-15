package cache

import (
	"context"

	"github.com/shal/hsa-2020/04/pkg/model"
)

type TransactionCache interface {
	Set(ctx context.Context, result *model.Result) error
	Get(ctx context.Context) (*model.Result, error)
}
