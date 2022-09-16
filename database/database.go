package database

import (
	"context"
	"fmt"
	"log"

	"github.com/MohammadMobasher/resturan-backend/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB(ctx context.Context, config models.Configuration) *mongo.Database {

	var cred options.Credential

	cred.Username = config.MongoUsername
	cred.Password = config.MongoPassword

	clientOption := options.Client().ApplyURI(config.MongoServer).SetAuth(cred)
	client, err := mongo.Connect(ctx, clientOption)
	if err != nil {
		panic(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		fmt.Println("Ping has problem")
		log.Fatal(err)
	}

	return client.Database(config.MongoDatabase)

}
