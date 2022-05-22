package service

import (
	"dou_sheng/mapper"
)

var userList = map[string]int{}

func GetUserList() {
	mapper.GetUserList(&userList)
}

func DoFavoriteAction(token string, videoID int, opt int) bool {
	userID, ok := userList[token]
	if !ok {
		return false
	}
	if opt == 1 {
		return mapper.LikeVideo(userID, videoID)
	} else {
		return mapper.DislikeVideo(userID, videoID)
	}
}
