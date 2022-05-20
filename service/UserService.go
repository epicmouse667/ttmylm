package service

import "dou_sheng/mapper"

var userList = map[string]int{}

func GetUserList() {
	mapper.GetUserList(&userList)
}
