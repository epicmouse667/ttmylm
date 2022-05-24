package controller

import (
	"dou_sheng/pogo"
	"dou_sheng/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type VideoListResponse struct {
	Response
	VideoList []pogo.Video `json:"video_list"`
}

func FavoriteList(c *gin.Context) {

	var favoriteList = service.GetFavoriteList(c.Query("token"), c.Query("user_id"))
	if favoriteList == nil {
		c.JSON(http.StatusOK, VideoListResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  "can't get the favorite list",
			},
			VideoList: nil,
		})
	}
	fmt.Println("favorite list: ", *favoriteList)
	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: *favoriteList,
	})
}
