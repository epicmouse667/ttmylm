package service

import (
	"dou_sheng/mapper"
	"dou_sheng/pogo"
)

func AddComment(user_id int, token string, video_id int, comment_text string) (*pogo.Comment, bool) {
	_, ok := userList[token]
	if !ok {
		return nil, false
	} else {
		return mapper.AddComment(user_id, video_id, comment_text), true
	}
}

func DeleteComment(user_id int, token string, video_id int, comment_id int) (*pogo.Comment, bool) {
	_, ok := userList[token]
	if !ok {
		return nil, false
	} else {
		return mapper.DeleteComment(user_id, video_id, comment_id), true
	}
}

func CommentList(token string, video_id int) (*[]pogo.Comment, bool) {
	_, ok := userList[token]
	if !ok {
		return nil, false
	} else {
		return mapper.GetCommentList(video_id), true
	}
}
