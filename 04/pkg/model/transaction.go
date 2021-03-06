package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Transaction struct {
	ID     primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Amount float64            `json:"amount" bson:"amount"`
	Time   time.Time          `json:"time" bson:"time"`
}

type Result struct {
	TotalCount int64   `json:"total_count"`
	TotalSum   float64 `json:"total_sum"`
}
