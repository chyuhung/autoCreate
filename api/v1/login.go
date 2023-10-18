package v1

import (
	"autoCreate/middleware"
	"autoCreate/model"
	"autoCreate/utils/errmsg"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) {
	var data model.User
	var token string
	c.ShouldBindJSON(&data)

	code := model.CheckLogin(data.Username, data.Password)
	if code == errmsg.SUCCSE {
		token, code = middleware.SetToken(data.Username)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
		"token":   token,
	})
}
