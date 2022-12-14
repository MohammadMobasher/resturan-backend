package mysql_database

import (
	"database/sql"
	"fmt"
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

func (f *FoodGroupMySqlRepository) GetItem(foodgroupId int64) (models.FoodGroupMySql, error) {
	var foodGroup models.FoodGroupMySql

	q := "SELECT * FROM food_group WHERE ID = ?"

	getItem, err := f.db.Prepare(q)

	if err != nil {
		log.Println(err)
		return foodGroup, err
	}

	err = getItem.QueryRow(foodgroupId).Scan(&foodGroup.Id, &foodGroup.Name, &foodGroup.ImageAddress)

	if err != nil {
		log.Println(err)
		return foodGroup, err
	}

	return foodGroup, nil

}

func (f *FoodGroupMySqlRepository) GetAll(skip int, take int) ([]models.FoodGroupMySql, int, error) {
	q := "SELECT * FROM food_group LIMIT " + fmt.Sprint(take) + " OFFSET " + fmt.Sprint(skip)
	countQuery := "SELECT count(*) FROM food_group"
	var count int
	items, err := f.db.Query(q)

	if err != nil {
		panic(err.Error())
	}

	err = f.db.QueryRow(countQuery).Scan(&count)
	if err != nil {
		panic(err.Error())
	}

	defer items.Close()

	var finalResult = []models.FoodGroupMySql{}

	for items.Next() {
		var foodGroup models.FoodGroupMySql
		err = items.Scan(&foodGroup.Id, &foodGroup.Name, &foodGroup.ImageAddress)
		if err != nil {
			return nil, 0, err
		}
		finalResult = append(finalResult, foodGroup)
	}

	return finalResult, count, nil

}

func (f *FoodGroupMySqlRepository) Insert(foodGroup models.FoodGroupMySql) (*models.FoodGroupMySql, error) {
	q := "INSERT INTO food_group(Name, ImageAddress) VALUES(?, ?)"
	insert, err := f.db.Prepare(q)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	resp, err := insert.Exec(foodGroup.Name, foodGroup.ImageAddress)
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

func (f *FoodGroupMySqlRepository) Delete(foodGroupId int) (bool, error) {
	q := "DELETE FROM food_group WHERE Id = ?"

	delete, err := f.db.Prepare(q)
	if err != nil {
		log.Println(err)
		return false, err
	}

	_, err = delete.Exec(foodGroupId)
	delete.Close()

	if err != nil {
		log.Println(err)
		return false, err
	}

	return true, nil

}

func (f *FoodGroupMySqlRepository) Update(obj models.FoodGroupMySql) (models.FoodGroupMySql, error) {
	var foodGroup models.FoodGroupMySql
	var q string
	if obj.ImageAddress != nil {
		q = "UPDATE food_group SET Name = ?, ImageAddress = ? WHERE Id = ?"
	} else {
		q = "UPDATE food_group SET Name = ? WHERE Id = ?"
	}

	update, err := f.db.Prepare(q)
	if err != nil {
		log.Println(err)
		return foodGroup, err
	}

	if obj.ImageAddress != nil {
		_, err = update.Exec(obj.Name, obj.ImageAddress, obj.Id)
	} else {
		_, err = update.Exec(obj.Name, obj.Id)
	}
	update.Close()

	if err != nil {
		log.Println(err)
		return foodGroup, err
	}

	return f.GetItem(obj.Id)
}
