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

	dbConf := config.Mongo
	dbName := dbConf.Dbname
	dbUser := dbConf.Username
	dbPass := dbConf.Password
	dbHost := dbConf.Host
	dbPort := dbConf.Port
	connectionString := fmt.Sprintf("mongodb://%s:%s@%s:%d/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	client, err := mongo.NewClient(connectionString)
	if err != nil {
		return nil, err
	}

	err = client.Connect(context.Background())
	if err != nil {
		return nil, err
	}

	// Collection types can be used to access the database
	db := client.Database(dbName)
	return db, nil
}
