package controller

import (
	"dou_sheng/pogo"
	"dou_sheng/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type FeedResponse struct {
	response  pogo.Response
	VideoList []pogo.Video `json:"video_list,omitempty"`
	NextTime  int64        `json:"next_time,omitempty"`
}

func Feed(c *gin.Context) {
	fmt.Println("feed.")
	list := service.GetFeedList(c.Query("token"))
	fmt.Println("feedList: ", list)
	if list == nil {
		c.JSON(http.StatusBadRequest, FeedResponse{
			response:  pogo.Response{StatusCode: 1},
			VideoList: nil,
			NextTime:  time.Now().Unix(),
		})
		return
	}
	c.JSON(http.StatusOK, FeedResponse{
		response:  pogo.Response{StatusCode: 0},
		VideoList: *list,
		NextTime:  time.Now().Unix(),
	})
}
