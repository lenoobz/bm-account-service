// Package service provides account service API
package service

import (
	"fmt"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/letrong/bm-account-service/conf"
	"github.com/letrong/bm-account-service/db"
	"github.com/mongodb/mongo-go-driver/mongo"
)

var config *conf.Config
var logger *logrus.Entry
var mongoDb *mongo.Database

func setSharedVariables(conf *conf.Config, log *logrus.Entry, mongo *mongo.Database) {
	config = conf
	logger = log
	mongoDb = mongo
}

// StartWebServer the entrypoint of web server
func StartWebServer(conf *conf.Config, log *logrus.Entry) {
	logger.Infoln("MongoDB connecting ...")
	mongo, err := db.ConnectMongoDB(conf, log)
	if err != nil {
		logger.Fatalln("Error occured while connecting to mongodb")
	}
	setSharedVariables(conf, log, mongo)

	r := NewRouter()
	http.Handle("/", r)

	port := fmt.Sprintf(":%d", config.Port)
	logger.Infof("Start web server on port %s. Running ...", port)
	err = http.ListenAndServe(port, nil)

	if err != nil {
		logger.Fatalf("Error occured while listening to port %d", config.Port)
		logger.WithError(err)
	}
}
