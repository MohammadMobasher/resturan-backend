package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

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
	Id          int64  `json:"id,omitempty" bson:"id,omitempty"`
	Name        string `json:"name,omitempty" bson:"name,omitempty" binding:"required"`
	FoodGroupId int64  `json:"FoodGroupId"  bson:"FoodGroupId,omitempty" binding:"required"`
	ResturanId  int64  `json:"resturanId"  bson:"resturanId,omitempty" binding:"required"`
}

type FoodImageMySql struct {
	Id           int64   `json:"id,omitempty" bson:"id,omitempty"`
	FoodId       int64   `json:"FoodGroupId"  bson:"FoodGroupId,omitempty" binding:"required"`
	ImageAddress *string `binding:"required"`
}

type FoodMySqlDTO struct {
	Id          int64          `json:"id,omitempty" bson:"id,omitempty"`
	Name        string         `json:"name,omzitempty" bson:"name,omitempty" binding:"required"`
	FoodGroupId int64          `json:"foodgroupid"  bson:"foodgroupid,omitempty" binding:"required"`
	ResturanId  int64          `json:"resturanid"  bson:"resturanid,omitempty"`
	Image       *string        `json:"image"  bson:"image"`
	Images      []string       `json:"images"  bson:"images"`
	FoodGroup   FoodGroupMySql `json:"foodgroup" bson:"foodgroup"`
	Resturan    ResturanMySql  `json:"resturan" bson:"resturan"`
}
