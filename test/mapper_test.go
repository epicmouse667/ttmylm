package test

import (
	"dou_sheng/mapper"
	"dou_sheng/util"
	"fmt"
	"testing"
)

func TestGetCommentList(t *testing.T) {
	util.InitSQL()
	commentlistptr := mapper.GetCommentList(2)
	fmt.Println(*commentlistptr)
}

func TestAddComment(t *testing.T) {
	util.InitSQL()
	commentptr := mapper.AddComment(1, 1, "lmao")
	fmt.Println(*commentptr)
}

func TestDeleteComment(t *testing.T) {
	util.InitSQL()
	commentptr := mapper.DeleteComment(1, 1, 78022)
	fmt.Println(*commentptr)
}

func TestGetFollowListByUserId(t *testing.T) {
	util.InitSQL()
	followlistptr := mapper.GetFollowListByUserId(1)
	fmt.Println(*followlistptr)
}

func TestFollowUser(t *testing.T) {
	util.InitSQL()
	err_msg_ptr := mapper.FollowUser(1, 3)
	if err_msg_ptr != nil {
		fmt.Print(*err_msg_ptr)
	}

}

func TestUnFollowUser(t *testing.T) {
	util.InitSQL()
	err_msg_ptr := mapper.UnFollowUser(1, 3)
	if err_msg_ptr != nil {
		fmt.Print(*err_msg_ptr)
	}

}
