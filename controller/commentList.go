package controller

import (
	"dou_sheng/pogo"
	"dou_sheng/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//extra_first 接口 4
type CommentListResponse struct {
	Response    Response
	CommentList []pogo.Comment
}

func CommentList(c *gin.Context) {
	token := c.Query("token")
	video_id, _ := strconv.Atoi(c.Query("video_id"))
	commentlist, flag := service.CommentList(token, video_id)
	if flag {
		c.JSON(http.StatusOK, CommentListResponse{
			Response: Response{
				StatusCode: 0,
				StatusMsg:  "Success",
			}, CommentList: *commentlist,
		})
	} else {
		c.JSON(http.StatusOK,
			CommentListResponse{
				Response: Response{
					StatusCode: 1,
					StatusMsg:  "Failed to fetch comment list.",
				}})

	}
}
