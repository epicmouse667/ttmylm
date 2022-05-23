package controller

import (
	"dou_sheng/mapper"
	"dou_sheng/pogo"
	"github.com/gin-gonic/gin"
	"net/http"
)

var userList = map[string]int{} //用户列表

type UserResponse struct {
	pogo.Response
	User pogo.User `json:"user"`
}

func UserInfo(c *gin.Context) {
	token := c.Query("token")
	user := mapper.GetUserByID(userList[token])
	c.JSON(http.StatusOK, UserResponse{
		Response: pogo.Response{StatusCode: 0},
		User:     user})

}
