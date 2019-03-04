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

// StartWebServer the entrypoint of web server
func StartWebServer(conf *conf.Config, log *logrus.Entry) {
	config = conf
	logger = log

	var err error
	mongoDb, err = db.ConnectMongoDB(config, logger)
	if err != nil {
		logger.Fatalln("Cannot connect to mongo")
	} else {
		logger.Infoln("MongoDB connected")

		r := NewRouter()
		http.Handle("/", r)

		port := fmt.Sprintf(":%d", config.Port)
		logger.Infof("Web server running ...")
		logger.Infof("Listen to port%s", port)
		err = http.ListenAndServe(port, nil)

		if err != nil {
			logger.Errorf("An error occured starting HTTP listener at port %d", config.Port)
			logger.Errorf("Error: " + err.Error())
		}
	}
}
