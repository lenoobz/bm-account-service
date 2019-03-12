// Package user provides abstraction layer to User core business logic
package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Control concrete type for user control
type Control struct {
	service Servicer
}

// InitUserControl initializes user control
func InitUserControl(service Servicer) *Control {
	return &Control{
		service,
	}
}

// CreateUser implements CreateUser contract of User Controller interface
func (uc *Control) CreateUser(c *gin.Context) {
	var u User
	c.BindJSON(&u)
	c.JSON(http.StatusOK, u)
}

// GetUsers implements GetUsers contract of User Controller interface
func (uc *Control) GetUsers(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong pong",
	})
}

// GetToken implements GetToken contract of User Controller interface
func (uc *Control) GetToken(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
