package mysql_database

import (
	"database/sql"
	"fmt"
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
	q := "INSERT INTO food(Name, FoodGroupId) VALUES(?, ?)"
	insert, err := f.db.Prepare(q)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	resp, err := insert.Exec(food.Name, food.FoodGroupId)
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

func (f *FoodMySqlRepository) GetAll(skip int, take int) ([]models.FoodMySqlDTO, int, error) {
	q := `SELECT f1.Id, f1.Name, f1.FoodGroupId, fi1.ImageAddress FROM food as f1
			LEFT JOIN food_image as fi1 ON fi1.Id = (select Id from food_image as fi2 where f1.Id = fi2.FoodId limit 1)
			LIMIT ` + fmt.Sprint(take) + " OFFSET " + fmt.Sprint(skip)
	// q := "SELECT * FROM food LIMIT " + fmt.Sprint(take) + " OFFSET " + fmt.Sprint(skip)
	// imageQuery := "SELECT * FROM food_image WHERE FoodId = ? ORDER BY Id LIMIT 1"
	countQuery := "SELECT count(*) FROM food"
	var count int

	items, err := f.db.Query(q)

	if err != nil {
		panic(err.Error())
	}

	err = f.db.QueryRow(countQuery).Scan(&count)
	if err != nil {
		panic(err.Error())
	}

	// images, err := f.db.Query(imageQuery)

	if err != nil {
		panic(err.Error())
	}

	defer items.Close()
	// defer images.Close()

	var finalResult = []models.FoodMySqlDTO{}

	for items.Next() {
		var food models.FoodMySqlDTO
		var image *string
		err = items.Scan(&food.Id, &food.Name, &food.FoodGroupId, &image)
		if err != nil {
			return nil, 0, err
		}
		// food.Images = append(food.Images, image)
		food.Image = image
		finalResult = append(finalResult, food)
	}

	return finalResult, count, nil

}

func (f *FoodMySqlRepository) GetItem(foodId int64) (models.FoodMySqlDTO, error) {
	var food models.FoodMySqlDTO

	q := ` SELECT 
				food.Id,
				food.Name,
				food.FoodGroupId,
				food.ResturanId,
				food_group.Id,
				food_group.Name,
				food_group.ImageAddress,
				resturan.Id,
				resturan.Name,
				resturan.Description
	 		from food 
			INNER JOIN food_group on food.FoodGroupId = food_group.Id 
			INNER JOIN resturan on food.ResturanId = resturan.Id
			WHERE food.Id = ? `
	q_images := `SELECT food_image.ImageAddress FROM food_image WHERE food_image.FoodId = ` + fmt.Sprint(foodId)
	q_comments := `SELECT * FROM food_comment WHERE food_comment.FoodId = ? `
	// q := "SELECT * FROM food LIMIT " + fmt.Sprint(take) + " OFFSET " + fmt.Sprint(skip)
	// imageQuery := "SELECT * FROM food_image WHERE FoodId = ? ORDER BY Id LIMIT 1"
	// q := "SELECT * FROM food_group WHERE ID = ?"

	getItem, err := f.db.Prepare(q)
	images, err := f.db.Query(q_images)
	comments, err := f.db.Query(q_comments, foodId)

	if err != nil {
		log.Println(err)
		return food, err
	}

	err = getItem.QueryRow(foodId).Scan(
		&food.Id,
		&food.Name,
		&food.FoodGroupId,
		&food.ResturanId,
		&food.FoodGroup.Id,
		&food.FoodGroup.Name,
		&food.FoodGroup.ImageAddress,
		&food.Resturan.Id,
		&food.Resturan.Name,
		&food.Resturan.Description)
	getItem.Close()

	for images.Next() {
		var image string
		err = images.Scan(&image)
		if err != nil {
			return food, err
		}
		food.Images = append(food.Images, image)
	}

	if err != nil {
		log.Println(err)
		return food, err
	}

	for comments.Next() {
		var comment models.FoodCommentMySql
		err = comments.Scan(&comment.Id, &comment.FoodId, &comment.Comment)
		if err != nil {
			return food, err
		}
		food.Comments = append(food.Comments, comment)
	}

	return food, nil

}

func (f *FoodMySqlRepository) CreateComment(foodComment models.FoodCommentMySql) (bool, error) {
	q := "INSERT INTO food_comment(FoodId, Comment) VALUES(?, ?)"

	insert, err := f.db.Prepare(q)
	if err != nil {
		log.Println(err)
		return false, err
	}

	_, err = insert.Exec(foodComment.FoodId, foodComment.Comment)
	insert.Close()

	if err != nil {
		log.Println(err)
		return false, err
	}

	return true, nil

}

func (f *FoodMySqlRepository) GetComments(foodId int64, skip int, take int) ([]models.FoodCommentMySql, int, error) {
	log.Println(skip)
	log.Println(take)
	q := "SELECT * FROM food_comment WHERE FoodId = ? LIMIT ? OFFSET ? "
	var finalResult = []models.FoodCommentMySql{}

	countQuery := "SELECT count(*) FROM food_comment where FoodId = ? "
	var count int

	err := f.db.QueryRow(countQuery, foodId).Scan(&count)
	if err != nil {
		panic(err.Error())
	}

	insert, err := f.db.Prepare(q)
	if err != nil {
		log.Println(err)
		return nil, 0, err
	}

	items, err := insert.Query(foodId, take, skip)
	insert.Close()

	if err != nil {
		log.Println(err)
		return nil, 0, err
	}

	for items.Next() {
		var item models.FoodCommentMySql
		err = items.Scan(&item.Id, &item.FoodId, &item.Comment)
		if err != nil {
			return nil, 0, err
		}
		finalResult = append(finalResult, item)
	}

	if err != nil {
		log.Println(err)
		return nil, 0, err
	}

	return finalResult, count, nil

}
