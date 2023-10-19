package v1

import (
	"autoCreate/middleware"
	"autoCreate/models"
	"autoCreate/utils/errmsg"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) {
	var data models.User
	var token string
	c.ShouldBindJSON(&data)

	code := models.CheckLogin(data.Username, data.Password)
	if code == errmsg.SUCCSE {
		token, code = middleware.SetToken(data.Username)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
		"token":   token,
	})
}
