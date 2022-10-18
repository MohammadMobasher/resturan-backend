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

func (f *FoodMySqlRepository) Insert(food models.FoodMySql, images []string) (*models.FoodMySql, error) {
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

	if len(images) > 0 {

		q_images := "INSERT INTO food_image(FoodId, ImageAddress) VALUES "
		vals := []interface{}{}

		for _, image := range images {
			q_images += "(?, ?),"
			vals = append(vals, food.Id, image)
		}

		log.Println(q_images)

		q_images = q_images[0 : len(q_images)-1]
		//prepare the statement
		insertImage, err := f.db.Prepare(q_images)

		if err != nil {
			log.Println(err)
			f.Delete(lastInsertId)
			return nil, err
		}

		//format all vals at once
		_, err = insertImage.Exec(vals...)
		if err != nil {
			log.Println(err)
			f.Delete(lastInsertId)
			return nil, err
		}
	}

	return &food, nil
}

func (f *FoodMySqlRepository) Delete(id int64) (bool, error) {
	q := "DELETE FROM food WHERE Id = ?"

	delete, err := f.db.Prepare(q)
	if err != nil {
		log.Println(err)
		return false, err
	}

	_, err = delete.Exec(id)
	delete.Close()

	if err != nil {
		log.Println(err)
		return false, err
	}

	return true, nil

}
