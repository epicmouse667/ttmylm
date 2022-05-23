package service

import (
	mapper "dou_sheng/mapper"
	"dou_sheng/pogo"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	//"sync/atomic"
)

var userList = map[string]int{} //用户列表

func GetUserList() {
	mapper.GetUserList(&userList)
}

//用户注册

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token := username + password

	if _, exist := userList[token]; exist {
		c.JSON(http.StatusOK, pogo.UserRegisterResponse{
			Response: pogo.Response{StatusCode: 1, StatusMsg: "User already exist"},
		})
	} else {
		userList[token] = len(userList)
		mapper.InsertUser(username, password)
		c.JSON(http.StatusOK, pogo.UserLoginResponse{
			Response: pogo.Response{StatusCode: 0, StatusMsg: "注册成功"},
			UserId:   len(userList),
			Token:    username + password,
		})
	}
}

//用户登录

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token := username + password
	user := userList[token]

	if mapper.SelectUser(username, password) {
		fmt.Println(userList)
		fmt.Println("succeed")
		c.JSON(http.StatusOK, pogo.UserLoginResponse{

			Response: pogo.Response{StatusCode: 0, StatusMsg: "成功登录"},
			UserId:   user,
			Token:    token,
		})
	} else {
		c.JSON(http.StatusOK, pogo.UserLoginResponse{
			Response: pogo.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}
