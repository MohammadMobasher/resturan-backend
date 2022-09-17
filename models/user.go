package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name     string             `json:"name,omitempty" bson:"name,omitempty" binding:"required"`
	UserName string             `json:"username,omitempty" binding:"required"`
	Password string             `json:"password,omitempty" binding:"required"`
}
