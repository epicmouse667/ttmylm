package controller

import (
	"dou_sheng/pogo"
)

//注册成功返回

type UserRegisterResponse struct {
	pogo.Response
	UserId int    `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

//登录成功返回

type UserLoginResponse struct {
	pogo.Response
	UserId int    `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

//用户注册

//用户登录
