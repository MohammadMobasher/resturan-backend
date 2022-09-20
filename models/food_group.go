package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type FoodGroup struct {
	Id   primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string             `json:"name,omitempty" bson:"name,omitempty" binding:"required"`
}
