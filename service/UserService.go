package service

import (
	"dou_sheng/mapper"
	"dou_sheng/pogo"
	"fmt"
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

func Register(username string, password string) (int, string) {
	token := username + password
	if mapper.FindUserByName(username) {
		return -1, ""
	}
	user := mapper.AddUser(username, password)
	if user == nil {
		fmt.Println("no user.")
		return -1, ""
	}
	userList[token] = user.Id
	return userList[token], token
}

//用户登录

func Login(username string, password string) (int, string) {

	token := username + password

	user := mapper.Login(username, password)

	if user == nil {
		fmt.Println("no user.")
		return -1, ""
	}
	return userList[token], token
}
func GetTokenInfo(token string) *pogo.User {
	id, ok := userList[token]
	if ok {
		return mapper.GetUserByID(id)
	} else {
		return nil
	}
}
