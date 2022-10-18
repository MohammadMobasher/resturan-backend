package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Food struct {
	Id           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name         string             `json:"name,omitempty" bson:"name,omitempty" binding:"required"`
	FoodGroupId  string             `json:"FoodGroupId"  bson:"FoodGroupId,omitempty" binding:"required"`
	ImageAddress string
}

type FoodImage struct {
	Id           primitive.ObjectID `json:"id,omitempty" bson:"id,omitempty"`
	FoodId       primitive.ObjectID `json:"FoodGroupId"  bson:"FoodGroupId,omitempty" binding:"required"`
	ImageAddress string             `binding:"required"`
}

///////////////////////////////////////////////////////////////////////

type FoodMySql struct {
	Id           int64  `json:"id,omitempty" bson:"id,omitempty"`
	Name         string `json:"name,omitempty" bson:"name,omitempty" binding:"required"`
	FoodGroupId  int64  `json:"FoodGroupId"  bson:"FoodGroupId,omitempty" binding:"required"`
	ImageAddress *string
}

type FoodImageMySql struct {
	Id           int64   `json:"id,omitempty" bson:"id,omitempty"`
	FoodId       int64   `json:"FoodGroupId"  bson:"FoodGroupId,omitempty" binding:"required"`
	ImageAddress *string `binding:"required"`
}
