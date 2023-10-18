package v1

// var code int

// // AddUser 添加用户
// func AddUser(c *gin.Context) {
// 	var data model.User
// 	if err := c.ShouldBindJSON(&data); err != nil {
// 		code = errmsg.ERROR
// 		c.JSON(http.StatusBadRequest, gin.H{"status": code, "message": err.Error()})
// 		return
// 	}

// 	// 验证
// 	msg, code := validator.Validate(&data)
// 	if code == errmsg.ERROR {
// 		c.JSON(http.StatusOK, gin.H{
// 			"status":  code,
// 			"message": msg,
// 		})
// 		return
// 	}

// 	code = model.CheckUser(data.Username)
// 	if code == errmsg.SUCCSE {
// 		model.CreateUser(&data)
// 	}
// 	if code == errmsg.ERROR_USERNAME_UESD {
// 		code = errmsg.ERROR_USERNAME_UESD
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"status":  code,
// 		"data":    data,
// 		"message": errmsg.GetErrMsg(code),
// 	})
// }
