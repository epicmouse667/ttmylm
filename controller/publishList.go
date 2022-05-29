package controller

import (
	"dou_sheng/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PublishList(c *gin.Context) {

	var publishList = service.GetPublishList(c.Query("token"), c.Query("user_id"))
	if publishList == nil {
		c.JSON(http.StatusOK, VideoListResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  "can't get the favorite list",
			},
			VideoList: nil,
		})
	}
	//fmt.Println("puublish list: ", *publishList)
	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: *publishList,
	})
}
