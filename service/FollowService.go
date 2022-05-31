package service

import (
	"dou_sheng/mapper"
	"dou_sheng/pogo"
)

func FollowUser(user_id int, token string, to_user_id int) (bool, *string) {
	if _, ok := userList[token]; !ok {
		err_msg := "Error:invalid user token"
		return false, &err_msg
	} else {
		err_msg := mapper.FollowUser(user_id, to_user_id)
		if err_msg == nil {
			//action_type==1  update followers
			mapper.UpdateFollow(user_id, to_user_id, 1)
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
		err_msg := mapper.UnFollowUser(user_id, to_user_id)
		if err_msg == nil {
			//actiontype==2 update followers
			mapper.UpdateFollow(user_id, to_user_id, 2)
			return true, err_msg
		} else {
			return false, err_msg
		}
	}
}

func GetFollowList(user_id int, token string) (bool, *string, *[]pogo.User) {
	if _, ok := userList[token]; !ok {
		err_msg := "Error:invalid user token"
		return false, &err_msg, nil
	} else {
		followlist := mapper.GetFollowListByUserId(user_id)
		return true, nil, followlist
	}
}
