package v1

import (
	"autoCreate/middleware"
	"autoCreate/models"
	"autoCreate/utils/errmsg"
	"net/http"

	"github.com/gin-gonic/gin"
)

// LoginHandler is a function to handle login requests
func LoginHandler(c *gin.Context) {
	// Create a variable to store the user data
	var data models.User
	// Create a variable to store the token
	var token string
	// Bind the JSON data to the data variable
	c.ShouldBindJSON(&data)

	// Call the CheckLogin function to check if the username and password are valid
	code := models.CheckLogin(data.Username, data.Password)
	// If the username and password are valid, set the token
	if code == errmsg.SUCCSE {
		token, code = middleware.SetToken(data.Username)
	}
	// Return the response with the status code, message, and token
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
		"token":   token,
	})
}
