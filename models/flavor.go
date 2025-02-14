package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Flavour struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	FlavourName string             `bson:"flavour_name" json:"flavour_name"`
	Quantity    string             `bson:"qty" json:"qty"`
}
