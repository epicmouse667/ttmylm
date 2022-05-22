package controller

import (
	"dou_sheng/pogo"
	"dou_sheng/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PublishList(c *gin.Context) {

	var publishList = service.GetPublishList(c.Query("token"), c.Query("user_id"))
	if publishList == nil {
		c.JSON(http.StatusOK, VideoListResponse{
			Response: pogo.Response{
				StatusCode: 1,
				StatusMsg:  "can't get the favorite list",
			},
			VideoList: nil,
		})
	}
	fmt.Println("puublish list: ", *publishList)
	c.JSON(http.StatusOK, VideoListResponse{
		Response: pogo.Response{
			StatusCode: 0,
		},
		VideoList: *publishList,
	})
}
