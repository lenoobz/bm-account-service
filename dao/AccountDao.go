// Package dao provides helper functions to query mongo database
package dao

import (
	"context"
	"log"

	"github.com/letrong/bm-account-service/model"
	"github.com/mongodb/mongo-go-driver/mongo"
)

// InsertOneValue inserts one item from Account model
func InsertOneValue(db *mongo.Database, account model.Account) {
	_, err := db.Collection("user").InsertOne(context.Background(), account)
	if err != nil {
		log.Fatal(err)
	}
}
