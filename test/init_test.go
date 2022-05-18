package test

import (
	"dou_sheng/mapper"
	"dou_sheng/pogo"
	"dou_sheng/util"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"testing"
	"time"
)

func TestInit(t *testing.T) {
	util.Init()
	func() {
		list := *mapper.GetFeedList()
		for i := 0; i < len(list); i++ {
			id := list[i].AuthorID
			list[i].Author = mapper.GetUserByID(id)
		}
		fmt.Println(list)
	}()
}

type FeedResponse struct {
	pogo.Response
	VideoList []pogo.Video `json:"video_list,omitempty"`
	NextTime  int64        `json:"next_time,omitempty"`
}

func TestRun(t *testing.T) {
	r := gin.Default()
	util.Init()
	r.Static("/static", "./public")

	apiRouter := r.Group("/douyin")

	// basic apis
	apiRouter.GET("/feed/", func(c *gin.Context) {
		list := *mapper.GetFeedList()
		for i := 0; i < len(list); i++ {
			id := list[i].AuthorID
			list[i].Author = mapper.GetUserByID(id)
		}
		fmt.Println(list)
		c.JSON(http.StatusOK, FeedResponse{
			Response:  pogo.Response{StatusCode: 0},
			VideoList: list,
			NextTime:  time.Now().Unix(),
		})
	})

	r.Run()
}
