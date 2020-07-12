package mongostore

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/shal/hsa-2020/04/pkg/model"
)

type TransactionRepository struct {
	collection *mongo.Collection
}

func NewTransactionRepository(db *mongo.Database) *TransactionRepository {
	return &TransactionRepository{
		collection: db.Collection("transactions"),
	}
}

func (r *TransactionRepository) Create(ctx context.Context, transaction *model.Transaction) error {
	res, err := r.collection.InsertOne(ctx, transaction)
	if err != nil {
		return err
	}

	// Assign ID of the inserted object.
	transaction.ID = res.InsertedID.(primitive.ObjectID)

	return nil
}

func (r *TransactionRepository) FindByID(ctx context.Context, id string) (*model.Transaction, error) {
	var transaction model.Transaction

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	err = r.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&transaction)
	if err != nil {
		return nil, err
	}

	return &transaction, nil
}

func (r *TransactionRepository) All(ctx context.Context) ([]model.Transaction, error) {
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	transactions := make([]model.Transaction, 0)
	if err := cursor.All(ctx, &transactions); err != nil {
		return nil, err
	}

	return transactions, nil
}
