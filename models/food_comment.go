package models

type FoodCommentMySql struct {
	Id      int64  `json:"id,omitempty" bson:"id,omitempty"`
	FoodId  int64  `json:"foodid,omitempty" bson:"foodid,omitempty" binding:"required"`
	Comment string `json:"comment" bson:"comment"`
	UserId  int64  `json:"userid" bson:"userid"`
}
