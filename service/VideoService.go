package service

import (
	"dou_sheng/mapper"
	"dou_sheng/pogo"
	"dou_sheng/util"
	"strconv"
)

var feedMap = map[string]*[]pogo.Video{}
var favoriteMap = map[int]*[]pogo.Video{}

func GetFeedList(token string) *[]pogo.Video {
	if len(userList) == 0 {
		GetUserList()
	}
	if token == "" {
		token = "_default_user_"
	}
	list, ok := feedMap[token]
	if ok {
		return list
	} else {
		list = mapper.GetFeedList(userList[token])
		util.Lock.Lock()
		feedMap[token] = list
		util.Lock.Unlock()
	}
	return list
}

func GetFavorateList(userID string, authorID string) *[]pogo.Video {
	userIDvar, _ := userList[userID]
	authorIDvar, _ := strconv.Atoi(authorID)
	return mapper.GetFavoriteList(userIDvar, authorIDvar)
}
