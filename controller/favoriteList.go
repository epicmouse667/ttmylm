package controller

import (
	"dou_sheng/pogo"
	"dou_sheng/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type VideoListResponse struct {
	pogo.Response
	VideoList []pogo.Video `json:"video_list"`
}

func FavoriteList(c *gin.Context) {

	var favoriteList = service.GetFavorateList(c.Query("token"), c.Query("user_id"))
	if favoriteList == nil {
		c.JSON(http.StatusOK, VideoListResponse{
			Response: pogo.Response{
				StatusCode: 1,
			},
			VideoList: nil,
		})
	}
	fmt.Println("favorite list: ", *favoriteList)
	c.JSON(http.StatusOK, VideoListResponse{
		Response: pogo.Response{
			StatusCode: 0,
		},
		VideoList: *favoriteList,
	})
}
