package mapper

import (
	"database/sql"
	"dou_sheng/pogo"
	"dou_sheng/util"
	"fmt"
	"log"
)

func GetFeedList(userID int) *[]pogo.Video {
	var list []pogo.Video
	stmt, err := util.DbConn.DB().Prepare("select user_id from user_favorite where user_id=? and video_id=?")
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	util.DbConn.Raw("select id, author_id,concat(url_pf,file_name) as play_url,concat(url_pf,cover_name) as cover_url,favorite_count,comment_count,title " +
		"from video " +
		"order by create_time desc").Scan(&list)
	for i := 0; i < len(list); i++ {
		id := list[i].AuthorID
		list[i].Author = GetUserByID(id)
		if GetUserRelation(userID, id) {
			list[i].Author.IsFollow = true
		} else {
			list[i].Author.IsFollow = false
		}
		t := -1
		err = stmt.QueryRow(userID, list[i].Id).Scan(&t)
		if err == sql.ErrNoRows {
			list[i].IsFavorite = false
		} else {
			list[i].IsFavorite = true
		}
		err = nil
	}
	fmt.Println(list)
	return &list
}
