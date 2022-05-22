package mapper

import (
	"dou_sheng/pogo"
	"dou_sheng/util"
)

func GetUserList(userList *map[string]int) {
	raw, _ := util.DbConn.DB().Query("select concat(name,password),id from user")
	var (
		key   string
		value int
	)
	for raw.Next() {
		raw.Scan(&key, &value)
		util.Lock.Lock()
		(*userList)[key] = value
		util.Lock.Unlock()
	}

}
func GetUserByID(id int) pogo.User {
	var user pogo.User
	util.DbConn.Raw("select id,name,follow_count,follower_count from user where id= ? ", id).Scan(&user)
	return user
}

func GetUserRelation(followerID int, userID int) bool {
	t := -1
	util.DbConn.DB().QueryRow("select user_id from user_follow where user_id=? and subscribe_id=?", followerID, userID).Scan(&t)
	return t != -1
}

func LikeVideo(userID int, videoID int) bool {
	util.DbConn.Lock()
	t := util.DbConn.Exec("insert into user_favorite values (?,?)", userID, videoID)
	if t.Error != nil {
		t.Error = nil
		//t.Rollback()
		util.DbConn.Unlock()
		return false
	}
	updateVideoFavoriteCount(true, videoID)
	t.Commit()
	util.DbConn.Unlock()
	return true
}
func DislikeVideo(userID int, videoID int) bool {
	util.DbConn.Lock()
	t := util.DbConn.Exec("delete from user_favorite where user_id= ? and video_id = ? ", userID, videoID)
	if t.Error != nil {
		t.Error = nil
		util.DbConn.Unlock()
		//t.Rollback()
		return false
	}
	updateVideoFavoriteCount(false, videoID)
	t.Commit()
	util.DbConn.Unlock()
	return true
}
