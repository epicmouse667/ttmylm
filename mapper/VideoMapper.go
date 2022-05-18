package mapper

import (
	"dou_sheng/pogo"
	"dou_sheng/util"
	"fmt"
)

var feedList pogo.Video

func GetFeedList() *[]pogo.Video {
	var list []pogo.Video
	util.DbConn.Raw("select id,author_id,play_url,cover_url,favorite_count,comment_count from video").Scan(&list)
	for i := 0; i < len(list); i++ {
		id := list[i].AuthorID
		list[i].Author = GetUserByID(id)
	}
	fmt.Println(list)
	return &list
}
