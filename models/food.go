package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Food struct {
	Id           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name         string             `json:"name,omitempty" bson:"name,omitempty" binding:"required"`
	FoodGroupId  string             `json:"FoodGroupId"  bson:"FoodGroupId,omitempty" binding:"required"`
	ImageAddress string
}
