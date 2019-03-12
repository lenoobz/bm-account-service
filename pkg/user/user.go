// Package user provides abstraction layer to User core business logic
package user

import "net/http"

// Repoer interface provide the abstraction layers wrap around user domain
type Repoer interface {
	CreateUser(user *User) error
	GetUser(username string) (*User, error)
}

// Servicer interface provide the abstraction layers wrap around user repo
type Servicer interface {
	CreateUser(user *User) error
	Login(username, password string) (*Login, error)
}

// Controller interface provide the abstraction layers wrap around user service
type Controller interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
	GetToken(w http.ResponseWriter, r *http.Request)
	//TODO Refresh Token
}
