package controller

import (
	"dou_sheng/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Feed(c *gin.Context) {
	fmt.Println("feed.")
	list := service.GetFeedList(c.Query("token"))
	fmt.Println("feedList: ", list)
	if list == nil {
		c.JSON(http.StatusBadRequest, FeedResponse{
			response: Response{
				StatusCode: 1,
				StatusMsg:  "cant' get the feed list",
			},
			VideoList: nil,
			NextTime:  time.Now().Unix(),
		})
		return
	}
	c.JSON(http.StatusOK, FeedResponse{
		response:  Response{StatusCode: 0},
		VideoList: *list,
		NextTime:  time.Now().Unix(),
	})
}
