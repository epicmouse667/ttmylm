package service

import (
	"dou_sheng/mapper"
	"dou_sheng/pogo"
	"dou_sheng/util"
	"fmt"
	"strconv"
)

var feedMap = map[string]int{}
var favoriteMap = map[int]*[]pogo.Video{}

func GetFeedList(token string) *[]pogo.Video {
	if len(userList) == 0 {
		GetUserList()
	}
	var list *[]pogo.Video
	if token == "" {
		token = "_default_user_"
	}
	index, ok := feedMap[token]
	if ok {
		util.Lock.Lock()
		list = mapper.GetFeedList(userList[token], index)
		feedMap[token]++
		util.Lock.Unlock()
	} else {
		util.Lock.Lock()
		list = mapper.GetFeedList(userList[token], 0)
		feedMap[token] = 1
		util.Lock.Unlock()
	}
	if list != nil && len(*list) == 0 {
		fmt.Println("service try again")
		util.Lock.Lock()
		feedMap[token] = 0
		util.Lock.Unlock()
		list = GetFeedList(token)
	}
	return list
}

func GetFavoriteList(userID string, authorID string) *[]pogo.Video {
	userIDvar, _ := userList[userID]
	authorIDvar, _ := strconv.Atoi(authorID)
	return mapper.GetFavoriteList(userIDvar, authorIDvar)
}
func GetPublishList(userID string, authorID string) *[]pogo.Video {
	userIDvar, _ := userList[userID]
	authorIDvar, _ := strconv.Atoi(authorID)
	return mapper.GetPublishList(userIDvar, authorIDvar)
}
