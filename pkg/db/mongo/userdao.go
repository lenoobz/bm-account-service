// Package dao provides abtraction layer to User databse
package dao

import (
	"github.com/letrong/bm-account-service/pkg/user"
	"github.com/mongodb/mongo-go-driver/mongo"
)

const tableName = "user"

// UserDao concrete type for user repo
type UserDao struct {
	db *mongo.Database
}

// InitUserDao initializes user dao
func InitUserDao(mongodb *mongo.Database) *UserDao {
	return &UserDao{
		mongodb,
	}
}

// CreateUser implements CreateUser contract of User Repoer interface
func (ud *UserDao) CreateUser(user *user.User) error {
	return nil
}

// GetUser implements GetUser contract of User Repoer interface
func (ud *UserDao) GetUser(username string) (*user.User, error) {
	return nil, nil
}
