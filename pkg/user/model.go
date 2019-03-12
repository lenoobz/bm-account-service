// Package user provides abstraction layer to User core business logic
package user

// User defines user model
type User struct {
	ID        string `json:"id,omitempty"`
	Username  string `json:"username"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Password  string `json:"password,omitempty"`
}

// Login defines user login model
type Login struct {
	Username string `json:"username"`
	Password string `json:"password,omitempty"`
	Token    string `json:"token"`
}
