package controller

import (
	"dou_sheng/pogo"
	"dou_sheng/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CommentActionResponse struct {
	Response Response
	Comment  pogo.Comment
}

func CommentAction(c *gin.Context) {
	user_id, _ := strconv.Atoi(c.Query("user_id"))
	token := c.Query("token")
	video_id, _ := strconv.Atoi(c.Query("video_id"))
	action_type, _ := strconv.Atoi(c.Query("action_type"))
	if action_type == 1 {
		comment_text := c.Query("comment_text")
		comment, flag := service.AddComment(user_id, token, video_id, comment_text)
		if flag {
			c.JSON(http.StatusOK,
				CommentActionResponse{
					Response: Response{
						StatusCode: 0,
						StatusMsg:  "Success",
					}, Comment: *comment})
		} else {
			c.JSON(http.StatusOK,
				CommentActionResponse{
					Response: Response{
						StatusCode: 1,
						StatusMsg:  "Failed to get comment list.",
					}})
		}
	} else {
		comment_id, _ := strconv.Atoi(c.Query("comment_id"))
		comment, flag := service.DeleteComment(user_id, token, video_id, comment_id)
		if flag {
			c.JSON(http.StatusOK,
				CommentActionResponse{
					Response: Response{
						StatusCode: 0,
						StatusMsg:  "Success",
					}, Comment: *comment})
		} else {
			c.JSON(http.StatusOK,
				CommentActionResponse{
					Response: Response{
						StatusCode: 1,
						StatusMsg:  "Failed to get comment list.",
					}})
		}
	}
}
