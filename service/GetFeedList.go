package service

import (
	"dou_sheng/pogo"
)
import "dou_sheng/mapper"

func GetFeedList() *[]pogo.Video {
	feedList := mapper.GetFeedList()
	return feedList
}
