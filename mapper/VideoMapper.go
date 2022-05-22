package mapper

import (
	"dou_sheng/pogo"
	"dou_sheng/util"
)

func fillListUser(userID int, list *pogo.Video) {
	id := list.AuthorID
	list.Author = GetUserByID(id)
	list.Author.IsFollow = GetUserRelation(userID, id)
}
func fillListRelation(userID int, list *pogo.Video) {
	t := -1
	util.Stmt.QueryRow(userID, list.Id).Scan(&t)
	list.IsFavorite = !(t == -1)
}

func GetFeedList(userID int) *[]pogo.Video {
	var list []pogo.Video
	util.DbConn.Raw("select id, author_id,concat(url_pf,file_name) as play_url,concat(url_pf,cover_name) as cover_url,favorite_count,comment_count,title " +
		"from video " +
		"order by create_time desc " +
		"limit 30").Scan(&list)
	for i := 0; i < len(list); i++ {
		fillListUser(userID, &list[i])
		fillListRelation(userID, &list[i])
	}
	return &list
}
func GetFavoriteList(userID int, authorID int) *[]pogo.Video {
	var list []pogo.Video
	util.DbConn.Raw("select id, author_id,concat(url_pf,file_name) as play_url,concat(url_pf,cover_name) as cover_url,favorite_count,comment_count,title "+
		"from video "+
		"where id in ( "+
		"   select video_id "+
		"   from user_favorite "+
		"   where user_id= ?    )", authorID).Scan(&list)
	for i := 0; i < len(list); i++ {
		fillListUser(userID, &list[i])
		fillListRelation(userID, &list[i])
	}
	return &list
}
func GetPublishList(userID int, authorID int) *[]pogo.Video {
	var list []pogo.Video
	util.DbConn.Raw("select id, author_id,concat(url_pf,file_name) as play_url,concat(url_pf,cover_name) as cover_url,favorite_count,comment_count,title "+
		"from video "+
		"where author_id = ? ", authorID).Scan(&list)
	for i := 0; i < len(list); i++ {
		fillListUser(userID, &list[i])
		fillListRelation(userID, &list[i])
	}
	return &list
}
func updateVideoFavoriteCount(increase bool, id int) {
	if increase {
		util.DbConn.Exec("update video set favorite_count=favorite_count+1 where id=?", id)
	} else {
		util.DbConn.Exec("update video set favorite_count=favorite_count-1 where id=?", id)
	}

}
