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
func GetUserByID(id int) *pogo.User {
	var user pogo.User
	util.DbConn.Raw("select id,name,follow_count,follower_count from user where id= ? ", id).Scan(&user)
	return &user
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
	//t.Commit()
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
	//t.Commit()
	util.DbConn.Unlock()
	return true
}
func AddUser(name string, password string) *pogo.User {
	util.DbConn.Lock()
	t := util.DbConn.Exec("insert into user(name, password, follow_count, follower_count) values (?,?,0,0)", name, password)
	if t.Error != nil {
		util.DbConn.Rollback()
		util.DbConn.Unlock()
		return nil
	}
	//util.DbConn.Commit()
	util.DbConn.Unlock()
	return Login(name, password)
}

func Login(name string, password string) *pogo.User {
	var user pogo.User
	t := util.DbConn.Raw("select * from user where name=? and password=?", name, password).Scan(&user)
	if t.Error != nil {
		return nil
	}
	return &user
}
func FindUserByName(name string) bool {
	t := -1
	util.DbConn.DB().QueryRow("select id from user where name=?", name).Scan(&t)
	return t != -1
}
