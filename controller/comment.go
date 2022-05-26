package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ttmylm/utils"
)

type CommentListRequest struct {
	VideoId string `json:"video_id"`
}
type CommentListResponse struct {
	Response
	CommentList []Comment `json:"comment_list,omitempty"`
}
type CommentActionRequest struct {
	UserId  string
	VideoId string
	Content string
}
type CommentActionResponse struct {
	Response
	Comment Comment `json:"comment,omitempty"`
}

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {
	var paramsRequest CommentActionRequest
	var paramsResponse CommentActionResponse
	if err := c.ShouldBindJSON(&paramsRequest); err != nil {
		paramsResponse.StatusCode = 200
		c.JSON(http.StatusOK, paramsResponse)
		return
	}

	token := c.Query("token")
	actionType := c.Query("action_type")

	if user, exist := usersLoginInfo[token]; exist {
		if actionType == "1" {
			text := c.Query("comment_text")
			c.JSON(http.StatusOK, CommentActionResponse{Response: Response{StatusCode: 0},
				Comment: Comment{
					Id:         1,
					User:       user,
					Content:    text,
					CreateDate: "05-01",
				}})
			return
		}
		c.JSON(http.StatusOK, Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {
	var listRequest CommentListRequest
	var listResponse CommentListResponse
	if err := c.ShouldBindJSON(&listRequest); err != nil {
		listResponse.StatusCode = 200
		c.JSON(http.StatusOK, listResponse)
		return
	}
	var context123 []Comment

	utils.Db.Table("comment").Find(&context123)
	listResponse.CommentList = context123
	c.JSON(http.StatusOK, listResponse)
	//c.JSON(http.StatusOK, CommentListResponse{
	//	Response:    Response{StatusCode: 0},
	//	CommentList: DemoComments,
	//})
}
