package controller

import (
	"dou_sheng/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func FavoriteAction(c *gin.Context) {
	opt, _ := strconv.Atoi(c.Query("action_type"))
	token := c.Query("token")
	videoID, _ := strconv.Atoi(c.Query("video_id"))
	res := service.DoFavoriteAction(token, videoID, opt)
	if res {
		fmt.Println("succeed")
		c.JSON(
			http.StatusOK,
			Response{
				StatusCode: 0,
			},
		)
	} else {
		fmt.Println("failed")
		c.JSON(
			http.StatusBadRequest,
			Response{
				StatusCode: 1,
				StatusMsg:  "err",
			},
		)
	}
}
