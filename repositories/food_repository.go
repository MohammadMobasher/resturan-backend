package repositories

import (
	"github.com/MohammadMobasher/resturan-backend/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IFoodRepository interface {
	// GetAll() ([]models.FoodGroup, error)
	Insert(foodGroup models.Food) models.Food
	// GetItem(foodGroupId string) (bool, error)
	// Delete(foodGroupId primitive.ObjectID) (bool, error)
	// Update(FoodGroupId models.FoodGroup) (*mongo.UpdateResult, error)
}

type FoodRepository struct {
	db *MongoRepository[models.Food]
}

func NewFoodRepository() *FoodRepository {
	return &FoodRepository{
		db: NewMongoRepository[models.Food]("food"),
	}
}

func (u *FoodRepository) Insert(food models.Food) (*models.Food, error) {

	result, err := u.db.Insert(food)
	food.Id = result.InsertedID.(primitive.ObjectID)
	return &food, err
}
