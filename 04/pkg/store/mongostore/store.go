package mongostore

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/shal/hsa-2020/04/pkg/config"
	"github.com/shal/hsa-2020/04/pkg/store"
)

type Store struct {
	db *mongo.Database

	transactionRepository *TransactionRepository
}

func New(ctx context.Context, conf config.Store) (*Store, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(conf.URI()))
	if err != nil {
		return nil, err
	}

	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}

	return &Store{
		db: client.Database(conf.Database),
	}, nil
}

func (s *Store) Transaction() store.TransactionRepository {
	if s.transactionRepository == nil {
		s.transactionRepository = NewTransactionRepository(s.db)
	}

	return s.transactionRepository
}
