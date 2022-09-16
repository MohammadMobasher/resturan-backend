package database

import (
	"context"

	"github.com/MohammadMobasher/resturan-backend/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB(ctx context.Context, config models.Configuration) *mongo.Database {

	var cred options.Credential

	cred.Username = config.MongoUsername
	cred.Password = config.MongoPassword

	clientOption := options.Client().ApplyURI(config.MongoServer)
	client, err := mongo.Connect(ctx, clientOption)
	if err != nil {
		panic(err)
	}
	return client.Database(config.MongoDatabase)

}
