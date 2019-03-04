// Package db provides abstract interface for database connection
package db

import (
	"context"
	"fmt"

	"github.com/Sirupsen/logrus"
	"github.com/letrong/bm-account-service/conf"
	"github.com/mongodb/mongo-go-driver/mongo"
)

// ConnectMongoDB create database connection
func ConnectMongoDB(config *conf.Config, logger *logrus.Entry) (*mongo.Database, error) {
	logger.Infoln("MongoDB connecting ...")
	dbConf := config.Mongo
	connectionString := fmt.Sprintf("mongodb://%s:%s@%s:%d/%s", dbConf.Username, dbConf.Password, dbConf.Host, dbConf.Port, dbConf.Dbname)

	client, err := mongo.NewClient(connectionString)
	if err != nil {
		logger.Fatal(err)
		return nil, err
	}

	err = client.Connect(context.Background())
	if err != nil {
		logger.Fatal(err)
		return nil, err
	}

	// Collection types can be used to access the database
	return client.Database(dbConf.Dbname), nil
}
