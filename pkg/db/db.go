// Package db provides abstract interface for database connection
package db

import (
	"fmt"

	"github.com/letrong/bm-account-service/conf"
	"github.com/mongodb/mongo-go-driver/mongo"
)

// InitMongoClient create database connection
func InitMongoClient(dbConf *conf.MongoConfig) (*mongo.Client, error) {

	dbName := dbConf.Dbname
	dbUser := dbConf.Username
	dbPass := dbConf.Password
	dbHost := dbConf.Host
	dbPort := dbConf.Port
	cxnString := fmt.Sprintf("mongodb://%s:%s@%s:%d/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	return mongo.NewClient(cxnString)
}
