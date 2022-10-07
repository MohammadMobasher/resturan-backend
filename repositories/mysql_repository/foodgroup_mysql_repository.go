package mysql_database

import (
	"database/sql"
	"log"

	"github.com/MohammadMobasher/resturan-backend/config"
	"github.com/MohammadMobasher/resturan-backend/database"
	"github.com/MohammadMobasher/resturan-backend/models"
)

type FoodGroupMySqlRepository struct {
	db *sql.DB
}

func NewFoodGroupMySqlRepository() *FoodGroupMySqlRepository {
	conf := config.GetConfig()
	db := database.ConnectMySqlDB(conf)

	return &FoodGroupMySqlRepository{
		db: db,
	}
}

func (f *FoodGroupMySqlRepository) Insert(foodGroup models.FoodGroupMySql) (*models.FoodGroupMySql, error) {
	q := "INSERT INTO food_group(Name) VALUES(?)"
	insert, err := f.db.Prepare(q)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	resp, err := insert.Exec(foodGroup.Name)
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

	foodGroup.Id = lastInsertId

	return &foodGroup, nil
}
