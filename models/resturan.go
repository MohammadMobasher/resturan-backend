package models

type ResturanMySql struct {
	Id          int64  `json:"id,omitempty" bson:"id,omitempty"`
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
}
