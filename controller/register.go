package controller

import (
	"dou_sheng/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

//用户注册

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token := username + password

	id, token := service.Register(username, password)

	if id == -1 {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User already exist"},
		})
	}
	c.JSON(http.StatusOK, UserLoginResponse{
		Response: Response{StatusCode: 0, StatusMsg: "succeed"},
		UserId:   id,
		Token:    token,
	})
}
