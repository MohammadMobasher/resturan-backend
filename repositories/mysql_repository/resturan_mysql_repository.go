package mysql_database

import (
	"database/sql"
	"fmt"
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

func (f *ResturanMySqlRepository) GetAll(skip int, take int) ([]models.ResturanMySql, int, error) {
	q := `SELECT * FROM resturan LIMIT ` + fmt.Sprint(take) + " OFFSET " + fmt.Sprint(skip)
	countQuery := "SELECT count(*) FROM resturan"
	var count int

	items, err := f.db.Query(q)

	if err != nil {
		panic(err.Error())
	}

	err = f.db.QueryRow(countQuery).Scan(&count)
	if err != nil {
		panic(err.Error())
	}

	var finalResult = []models.ResturanMySql{}

	for items.Next() {
		var resturan models.ResturanMySql
		err = items.Scan(&resturan.Id, &resturan.Name, &resturan.Description)
		if err != nil {
			return nil, 0, err
		}
		finalResult = append(finalResult, resturan)
	}

	return finalResult, count, nil

}

func (f *ResturanMySqlRepository) Delete(resturanId int) (bool, error) {
	q := "DELETE FROM resturan WHERE Id = ?"

	delete, err := f.db.Prepare(q)
	if err != nil {
		log.Println(err)
		return false, err
	}

	_, err = delete.Exec(resturanId)
	delete.Close()

	if err != nil {
		log.Println(err)
		return false, err
	}

	return true, nil

}
