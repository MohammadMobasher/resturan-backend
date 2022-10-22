package mysql_database

import (
	"database/sql"
	"log"

	"github.com/MohammadMobasher/resturan-backend/config"
	"github.com/MohammadMobasher/resturan-backend/database"
	"github.com/MohammadMobasher/resturan-backend/models"
)

// type IResturanRepository interface {
// 	// GetAll() ([]models.FoodGroup, error)
// 	Insert(foodGroup models.Food) models.Food
// 	// GetItem(foodGroupId string) (bool, error)
// 	// Delete(foodGroupId primitive.ObjectID) (bool, error)
// 	// Update(FoodGroupId models.FoodGroup) (*mongo.UpdateResult, error)
// }

type ResturanMySqlRepository struct {
	db *sql.DB
}

func NewResturanMySqlRepository() *ResturanMySqlRepository {
	conf := config.GetConfig()
	db := database.ConnectMySqlDB(conf)

	return &ResturanMySqlRepository{
		db: db,
	}
}

func (f *ResturanMySqlRepository) Insert(resturan models.ResturanMySql) (*models.ResturanMySql, error) {
	q := "INSERT INTO resturan(Name, Description) VALUES(?, ?)"
	insert, err := f.db.Prepare(q)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	resp, err := insert.Exec(resturan.Name, resturan.Description)
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

	resturan.Id = lastInsertId

	return &resturan, nil
}
