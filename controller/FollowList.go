package controller

import (
	"dou_sheng/pogo"
	"dou_sheng/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FollowListResponse struct {
	Response Response
	UserList []pogo.User
}

func FollowList(c *gin.Context) {
	user_id, _ := strconv.Atoi(c.Query("user_id"))
	token := c.Query("token")
	flag, status_msg, user_list := service.GetFollowList(user_id, token)
	if flag {
		if status_msg == nil {
			c.JSON(http.StatusOK, FollowListResponse{
				Response: Response{
					StatusCode: 0,
					StatusMsg:  "Success.",
				}, UserList: *user_list,
			})
		} else {
			c.JSON(http.StatusOK, FollowListResponse{
				Response: Response{
					StatusCode: 1,
					StatusMsg:  *status_msg,
				},
			})
		}

	} else {
		c.JSON(http.StatusOK, FollowListResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  *status_msg,
			},
		})
	}
}
