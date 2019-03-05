// Package dao provides helper functions to query mongo database
package dao

import (
	"context"

	"github.com/Sirupsen/logrus"
	"github.com/letrong/bm-account-service/model"
	"github.com/mongodb/mongo-go-driver/mongo"
)

const accountCollection = "user"

// InsertNewAccount inserts new account
func InsertNewAccount(mongoDb *mongo.Database, logger *logrus.Entry, account model.Account) {
	_, err := mongoDb.Collection(accountCollection).InsertOne(context.Background(), account)
	if err != nil {
		logger.Fatalln(err)
	}
}
