package service

import (
	"dou_sheng/mapper"
	"dou_sheng/pogo"
	"fmt"
	"sync"
)

var feedMap = map[string]*[]pogo.Video{}

func GetFeedList(token string) *[]pogo.Video {
	if len(userList) == 0 {
		GetUserList()
		fmt.Println(userList)
	}
	if token == "" {
		token = "_default_user_"
	}
	fmt.Println("user: " + token)
	list, ok := feedMap[token]
	if ok {
		return list
	} else {
		list = mapper.GetFeedList(userList[token])
		var lock sync.Mutex
		lock.Lock()
		feedMap[token] = list
		lock.Unlock()
	}
	return list
}
