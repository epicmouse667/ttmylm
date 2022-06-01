package test

import (
	"dou_sheng/service"
	"dou_sheng/util"
	"fmt"
	"testing"
)

func TestAddCommentService(t *testing.T) {
	util.InitSQL()
	service.GetUserList()
	comment_ptr, _ := service.AddComment(1, "zhangleidouyin", 1, "lol")
	fmt.Println(*comment_ptr)
}

func TestDeleteCommentService(t *testing.T) {
	util.InitSQL()
	service.GetUserList()
	comment_ptr, _ := service.DeleteComment(1, "zhangleidouyin", 1, 78021)
	fmt.Println(*comment_ptr)
}
