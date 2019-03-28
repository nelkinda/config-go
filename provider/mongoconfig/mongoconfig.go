package mongoconfig

import (
	"context"
	"fmt"
	"github.com/nelkinda/config-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type env struct {
	Environment   string
	Configuration map[string]string
}

// Configuration to setup a MongoConfigProvider
type Config struct {
	// URI to connect with MongoDB.
	// Defaults to "" which means connecting to localhost.
	URI string

	// Database from which to read the configuration, defaults to "config"
	Database string

	// Collection from which to read the configuration, defaults to "config"
	Collection string

	// Environment (key) in the configuration collection, defaults to "production"
	Environment string
}

func CreateMongoConfigProvider(c *Config) config.ProviderFunc {
	if c.Database == "" {
		c.Database = "config"
	}
	if c.Collection == "" {
		c.Collection = "config"
	}
	if c.Environment == "" {
		c.Environment = "production"
	}
	if client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(c.URI)); err != nil {
		panic(err)
	} else {
		database := client.Database(c.Database)
		collection := database.Collection(c.Collection)
		ms := bson.M{"environment": c.Environment}
		return func(key string) (string, error) {
			e := env{}
			if err := collection.FindOne(context.Background(), ms).Decode(&e); err != nil {
				return "", fmt.Errorf("configuration collection not found for URI %s database %s collection %s environment %s key %s", c.URI, c.Database, c.Collection, c.Environment, key)
			}
			if value, ok := e.Configuration[key]; ok {
				return value, nil
			} else {
				return "", fmt.Errorf("configuration value not found for URI %s database %s collection %s environment %s key %s", c.URI, c.Database, c.Collection, c.Environment, key)
			}
		}
	}
}
