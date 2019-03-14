package mongoconfig

import (
	"context"
	"github.com/nelkinda/config-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type entry struct {
	Key string
	Value string
}

func CreateMongoConfigProvider(mongodbUrl, databaseName, collectionName  string) config.ProviderFunc {
	if client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongodbUrl)); err != nil {
		panic(err)
	} else {
		database := client.Database(databaseName)
		collection := database.Collection(collectionName)
		return func(key string) (string, error) {
			e := entry{}
			if err:= collection.FindOne(context.Background(), bson.M{"key": key}).Decode(&e); err != nil {
				return "", err
			} else {
				return e.Value, nil
			}
		}
	}
}
