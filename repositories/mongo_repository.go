package repositories

import (
	"context"
	"log"

	"github.com/MohammadMobasher/resturan-backend/config"
	"github.com/MohammadMobasher/resturan-backend/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRepository[T any] struct {
	Ctx        context.Context
	Collection *mongo.Collection
}

func NewMongoRepository[T any](collectionName string) *MongoRepository[T] {
	conf := config.GetConfig()
	ctx := context.TODO()

	db := database.ConnectDB(ctx, conf)

	return &MongoRepository[T]{
		Ctx:        ctx,
		Collection: db.Collection(collectionName),
	}
}

func (m *MongoRepository[T]) Insert(obj T) (*mongo.InsertOneResult, error) {

	result, err := m.Collection.InsertOne(m.Ctx, obj)

	if err != nil {
		log.Fatal(err)
	}

	return result, err
}

func (m *MongoRepository[T]) Delete(filter interface{}) (*mongo.DeleteResult, error) {

	result, err := m.Collection.DeleteOne(m.Ctx, filter)

	if err != nil {
		log.Fatal(err)
	}

	return result, err
}

func (m *MongoRepository[T]) InsertMany(objs []T) (*mongo.InsertManyResult, error) {
	obj := make([]interface{}, len(objs))

	for i, v := range objs {
		obj[i] = v
	}
	result, err := m.Collection.InsertMany(m.Ctx, obj)

	if err != nil {
		log.Fatal(err)
	}

	return result, err
}

func (m *MongoRepository[T]) FindMany(filter interface{}, findOptions *options.FindOptions) ([]T, error) {
	var result []T
	cur, err := m.Collection.Find(m.Ctx, filter, findOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(m.Ctx)

	for cur.Next(m.Ctx) {
		var obj T
		err := cur.Decode(&obj)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, obj)
	}

	if err != nil {
		log.Fatal(err)
	}

	return result, err
}

func (m *MongoRepository[T]) Update(id primitive.ObjectID, obj T) (*mongo.UpdateResult, error) {
	filter := bson.M{"_id": id}
	log.Println(id)
	var returnResult T

	err1 := m.Collection.FindOne(m.Ctx, filter).Decode(&returnResult)
	if err1 == mongo.ErrNoDocuments {
		log.Println("record does not exist")
	} else if err1 != nil {
		log.Fatal(err1)
	}

	object, err := bson.Marshal(obj)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	var update bson.M
	err = bson.Unmarshal(object, &update)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	log.Println(update)
	log.Println(filter)

	result, err := m.Collection.UpdateOne(m.Ctx, filter, bson.D{{Key: "$set", Value: update}})

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return result, err

}

func (m *MongoRepository[T]) FindOne(filter interface{}) (T, error) {
	var finalResult T
	err := m.Collection.FindOne(m.Ctx, filter).Decode(&finalResult)

	if err == mongo.ErrNoDocuments {
		return finalResult, nil
	} else if err != nil {
		log.Fatal(err)
		return finalResult, err
	}

	return finalResult, nil

}
