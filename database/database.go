package database

import (
	"context"

	"github.com/MohammadMobasher/resturan-backend/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB(ctx context.Context, config models.Configuration) *mongo.Database {

	clientOption := options.Client().ApplyURI(config.MongoServer)
	client, err := mongo.Connect(ctx, clientOption)
	if err != nil {
		panic(err)
	}
	return client.Database(config.MongoDatabase)

}
