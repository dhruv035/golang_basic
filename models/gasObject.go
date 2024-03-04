package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type GasObject struct {
	ID        primitive.ObjectID `bson:"_id"`
	MedGas    int                `bson:"medGas"`
	Timestamp int                `bson:"timestamp"`
}
