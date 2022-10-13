package mysql_database

import (
	"database/sql"
	"log"

	"github.com/MohammadMobasher/resturan-backend/config"
	"github.com/MohammadMobasher/resturan-backend/database"
	"github.com/MohammadMobasher/resturan-backend/models"
)

type IFoodRepository interface {
	// GetAll() ([]models.FoodGroup, error)
	Insert(foodGroup models.Food) models.Food
	// GetItem(foodGroupId string) (bool, error)
	// Delete(foodGroupId primitive.ObjectID) (bool, error)
	// Update(FoodGroupId models.FoodGroup) (*mongo.UpdateResult, error)
}

type FoodMySqlRepository struct {
	db *sql.DB
}

func NewFoodMySqlRepository() *FoodMySqlRepository {
	conf := config.GetConfig()
	db := database.ConnectMySqlDB(conf)

	return &FoodMySqlRepository{
		db: db,
	}
}

func (f *FoodMySqlRepository) Insert(food models.FoodMySql) (*models.FoodMySql, error) {
	q := "INSERT INTO food(Name, FoodGroupId, ImageAddress) VALUES(?, ?, ?)"
	insert, err := f.db.Prepare(q)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	resp, err := insert.Exec(food.Name, food.FoodGroupId, food.ImageAddress)
	insert.Close()

	if err != nil {
		log.Println(err)
		return nil, err
	}

	lastInsertId, err := resp.LastInsertId()

	if err != nil {
		log.Println(err)
		return nil, err
	}

	food.Id = lastInsertId

	return &food, nil
}
