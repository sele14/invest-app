package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// model of our data
type Instrument struct {
	ID       primitive.ObjectID 	`bson:"_id"`
	Type     string  `json:"type"`
	Name     string  `json:"name"`
	Price    float32 `json:"price"`
	Quantity int     `json:"quantity"`
}
