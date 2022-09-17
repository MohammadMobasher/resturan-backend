package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name     string             `json:"name,omitempty"`
	UserName string             `json:"username,omitempty"`
	Password string             `json:"password,omitempty" bson:"password,omitempty"`
}
