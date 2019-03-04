// Package model provides data entity mapping for account service
package model

// Account model
type Account struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
