package mongo_repositories

import (
	"log"

	"github.com/MohammadMobasher/resturan-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type IUserRepository interface {
	GetAll() ([]models.User, error)
	Insert(user models.User) models.User
	GetItem(userId string) (bool, error)
	Delete(userId primitive.ObjectID) (bool, error)
	Update(user models.User) (*mongo.UpdateResult, error)
}

type UserRepository struct {
	db *MongoRepository[models.User]
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		db: NewMongoRepository[models.User]("user"),
	}
}

func (u *UserRepository) Insert(user models.User) (*models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	user.Password = string(hashedPassword)
	result, err := u.db.Insert(user)
	user.Id = result.InsertedID.(primitive.ObjectID)
	return &user, err
}

func (u *UserRepository) Delete(userId string) (bool, error) {
	id, _ := primitive.ObjectIDFromHex(userId)
	_, err := u.db.Delete(bson.M{"_id": id})

	if err != nil {
		return false, nil
	}

	return true, nil
}

func (u *UserRepository) GetItem(userId string) (models.User, error) {
	id, _ := primitive.ObjectIDFromHex(userId)
	return u.db.FindOne(bson.M{"_id": id})
}

func (u *UserRepository) GetAll() ([]models.User, error) {
	result, err := u.db.FindMany(bson.D{}, nil)
	return result, err
}

func (u *UserRepository) Update(user models.User) (*mongo.UpdateResult, error) {

	result, err := u.db.Update(user.Id, user)
	return result, err
}
