package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetClient() (*mongo.Client, *mongo.Database, error) {
	// Get MongoDB Config
	mongodbConfig := GetConfig()

	ctxRoot := context.Background()
	ctx, cancel := context.WithTimeout(ctxRoot, 10*time.Second)

	defer cancel()

	log.Println("MongoDB URL ~> " + mongodbConfig.DBAddress)

	// MongoDB Client
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongodbConfig.DBAddress))
	if err != nil {
		return nil, nil, err
	}

	// MongoDB DataBase
	db := client.Database(mongodbConfig.DBName)

	/*
		defer func() {
			if err = client.Disconnect(ctx); err != nil {
				panic(err)
			}
		}()
	*/
	return client, db, err
}
