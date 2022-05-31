package controller

import (
	"dou_sheng/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RelationAction(c *gin.Context) {
	var flag bool
	var status_msg *string
	user_id, _ := strconv.Atoi(c.Query("user_id"))
	token := c.Query("token")
	to_user_id, _ := strconv.Atoi(c.Query("to_user_id"))
	action_type, _ := strconv.Atoi(c.Query("action_type"))
	if action_type == 1 {
		flag, status_msg = service.FollowUser(user_id, token, to_user_id)
	} else {
		flag, status_msg = service.UnFollowUser(user_id, token, to_user_id)
	}
	if flag {
		if status_msg == nil {
			c.JSON(http.StatusOK, Response{
				StatusCode: 0,
				StatusMsg:  "Success",
			})
		} else {
			c.JSON(http.StatusOK, Response{
				StatusCode: 1,
				StatusMsg:  *status_msg,
			})
		}

	} else {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  *status_msg,
		})
	}

}
