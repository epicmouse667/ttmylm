package controller

import (
	"dou_sheng/pogo"
	"dou_sheng/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type PublishResponse struct {
	response pogo.Response
}

func Publish(c *gin.Context) {
	var form
	// Bind()默认解析并绑定form格式
	// 根据请求头中content-type自动推断
	if err := c.Bind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 判断用户名密码是否正确
	if form.User != "root" || form.Pssword != "admin" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "200"})
}
