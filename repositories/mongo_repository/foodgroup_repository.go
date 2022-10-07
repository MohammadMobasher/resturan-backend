package mongo_repositories

import (
	"github.com/MohammadMobasher/resturan-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type IFoodGroupRepository interface {
	GetAll() ([]models.FoodGroup, error)
	Insert(foodGroup models.FoodGroup) models.FoodGroup
	GetItem(foodGroupId string) (bool, error)
	Delete(foodGroupId primitive.ObjectID) (bool, error)
	Update(FoodGroupId models.FoodGroup) (*mongo.UpdateResult, error)
}

type FoodGroupRepository struct {
	db *MongoRepository[models.FoodGroup]
}

func NewFoodGroupRepository() *FoodGroupRepository {
	return &FoodGroupRepository{
		db: NewMongoRepository[models.FoodGroup]("foodgroup"),
	}
}

func (u *FoodGroupRepository) Insert(foodGroup models.FoodGroup) (*models.FoodGroup, error) {

	result, err := u.db.Insert(foodGroup)
	foodGroup.Id = result.InsertedID.(primitive.ObjectID)
	return &foodGroup, err
}

func (u *FoodGroupRepository) Delete(foodGroupId string) (bool, error) {
	id, _ := primitive.ObjectIDFromHex(foodGroupId)
	_, err := u.db.Delete(bson.M{"_id": id})

	if err != nil {
		return false, nil
	}

	return true, nil
}

func (u *FoodGroupRepository) GetItem(foodGroupId string) (models.FoodGroup, error) {
	id, _ := primitive.ObjectIDFromHex(foodGroupId)
	return u.db.FindOne(bson.M{"_id": id})
}

func (u *FoodGroupRepository) GetAll() ([]models.FoodGroup, error) {
	result, err := u.db.FindMany(bson.D{}, nil)
	return result, err
}

func (u *FoodGroupRepository) Update(user models.FoodGroup) (*mongo.UpdateResult, error) {

	result, err := u.db.Update(user.Id, user)
	return result, err
}
