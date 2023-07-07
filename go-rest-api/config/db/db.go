package db

import (
	"context"
	"errors"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = context.TODO()

func ConnectDB() (*mongo.Database, error) {

	clientOptions := options.Client().ApplyURI(os.Getenv("DB_URI"))

	client, dbErr := mongo.Connect(ctx, clientOptions)
	if dbErr != nil {
		return nil, errors.New("Failed to connect to database " + dbErr.Error())
	}

	dbErr = client.Ping(ctx, nil)
	if dbErr != nil {
		return nil, errors.New("Failed to connect to database " + dbErr.Error())
	}

	databaseName := os.Getenv("DB_NAME")
	if databaseName == "" {
		return nil, errors.New("Please provide the database name")
	}

	database := client.Database(databaseName)

	return database, nil
}
