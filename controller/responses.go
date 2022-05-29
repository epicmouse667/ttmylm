package controller

import "dou_sheng/pogo"

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}
type FeedResponse struct {
	response  Response
	VideoList []pogo.Video `json:"video_list,omitempty"`
	NextTime  int64        `json:"next_time,omitempty"`
}
type UserRegisterResponse struct {
	Response
	UserId int    `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserLoginResponse struct {
	Response
	UserId int    `json:"user_id,omitempty"`
	Token  string `json:"token"`
}
type UserResponse struct {
	Response
	User pogo.User `json:"user"`
}
