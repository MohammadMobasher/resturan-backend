package mysql_database

import (
	"database/sql"
	"log"

	"github.com/MohammadMobasher/resturan-backend/config"
	"github.com/MohammadMobasher/resturan-backend/database"
	"github.com/MohammadMobasher/resturan-backend/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type IUserMySqlRepository interface {
	GetAll() ([]models.UserMySql, error)
	Insert(user models.UserMySql) models.UserMySql
	GetItem(userId string) (bool, error)
	Delete(userId primitive.ObjectID) (bool, error)
	Update(user models.UserMySql) (*mongo.UpdateResult, error)
}

type UserMySqlRepository struct {
	db *sql.DB
}

func NewUserMySqlRepository() *UserMySqlRepository {
	conf := config.GetConfig()
	db := database.ConnectMySqlDB(conf)

	return &UserMySqlRepository{
		db: db,
	}
}

func (u *UserMySqlRepository) Insert(user models.UserMySql) (*models.UserMySql, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	user.Password = string(hashedPassword)

	q := "INSERT INTO user(Name, UserName, Password) VALUES(?, ?, ?)"

	insert, err := u.db.Prepare(q)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	resp, err := insert.Exec(user.Name, user.UserName, user.Password)
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

	user.Id = lastInsertId

	return &user, nil
}

func (u *UserMySqlRepository) Delete(userId string) (bool, error) {
	q := "DELETE FROM user WHERE Id = ?"

	delete, err := u.db.Prepare(q)
	if err != nil {
		log.Println(err)
		return false, err
	}

	_, err = delete.Exec(userId)
	delete.Close()

	if err != nil {
		log.Println(err)
		return false, err
	}

	return true, nil

}

func (u *UserMySqlRepository) GetItem(userId int64) (models.UserMySql, error) {
	var user models.UserMySql

	q := "SELECT * FROM user WHERE Id = ?"

	getItem, err := u.db.Prepare(q)

	if err != nil {
		log.Println(err)
		return user, err
	}

	err = getItem.QueryRow(userId).Scan(&user.Id, &user.Name)

	if err != nil {
		log.Println(err)
		return user, err
	}

	return user, nil
}

func (u *UserMySqlRepository) GetAll() ([]models.UserMySql, error) {
	q := "SELECT * FROM user"

	items, err := u.db.Query(q)

	if err != nil {
		panic(err.Error())
	}

	defer items.Close()

	var finalResult = []models.UserMySql{}

	for items.Next() {
		var user models.UserMySql
		err = items.Scan(&user.Id, &user.Name)
		if err != nil {
			return nil, err
		}
		finalResult = append(finalResult, user)
	}

	return finalResult, nil
}

func (u *UserMySqlRepository) Update(obj models.UserMySql) (models.UserMySql, error) {
	var user models.UserMySql
	q := "UPDATE user SET Name = ?, UserName = ? WHERE Id = ?"

	update, err := u.db.Prepare(q)
	if err != nil {
		log.Println(err)
		return user, err
	}

	_, err = update.Exec(obj.Name, obj.UserName, obj.Id)
	update.Close()

	if err != nil {
		log.Println(err)
		return user, err
	}

	return u.GetItem(obj.Id)
}
