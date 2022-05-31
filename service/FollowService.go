package service

import "dou_sheng/mapper"

func FollowUser(user_id int, token string, to_user_id int) (bool, *string) {
	if _, ok := userList[token]; !ok {
		err_msg := "Error:invalid user token"
		return false, &err_msg
	} else {
		err_msg := mapper.FollowUser(user_id, to_user_id)
		if err_msg == nil {
			return true, err_msg
		} else {
			return false, err_msg
		}
	}
}

func UnFollowUser(user_id int, token string, to_user_id int) (bool, *string) {
	if _, ok := userList[token]; !ok {
		err_msg := "Error:invalid user token"
		return false, &err_msg
	} else {
		err_msg := mapper.FollowUser(user_id, to_user_id)
		if err_msg == nil {
			return true, err_msg
		} else {
			return false, err_msg
		}
	}
}

func GetFollowList() {}
