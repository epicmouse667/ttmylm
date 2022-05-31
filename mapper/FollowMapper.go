package mapper

import (
	"dou_sheng/pogo"
	"dou_sheng/util"
	"log"
)

func GetFollowListByUserId(user_id int) *[]pogo.User {
	//followlist full capcity:1000
	followlist := make([]pogo.User, 0, 1000)
	rows, err := util.DbConn.Raw(`
	select subsribe_id
	from followlist
	where user_id=?
	`, user_id).Rows()
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var follower pogo.User
		rows.Scan(&follower.Id, follower.Name, follower.FollowCount, follower.FollowerCount)
		follower.IsFollow = true
		followlist = append(followlist, follower)
	}
	return &followlist
}

func FollowUser(user_id int, to_user_id int) *string {
	if GetUserRelation(user_id, to_user_id) {
		err_msg := "Error:you have alreday followed this user"
		return &err_msg
	}
	util.DbConn.Lock()
	t := util.DbConn.Exec(
		`insert into
		user_follow 
		values(?,?)`, user_id, to_user_id)
	if t.Error != nil {
		util.DbConn.Rollback()
		util.DbConn.Unlock()
	}
	util.DbConn.Unlock()
	return nil
}

func UnFollowUser(user_id int, to_user_id int) *string {
	if !GetUserRelation(user_id, to_user_id) {
		err_msg := "Error:you haven't followed this user yet"
		return &err_msg
	}
	util.DbConn.Lock()
	t := util.DbConn.Exec(
		`delete form
			user_follow 
		where
			user_id=? AND subscribe_id=?`, user_id, to_user_id)
	if t.Error != nil {
		util.DbConn.Rollback()
		util.DbConn.Unlock()
	}
	util.DbConn.Unlock()
	return nil
}

func UpdateFollow(user_id int, to_user_id int, action_type int) {
	util.DbConn.Lock()
	//action_type==1 then follower+1 offset=1
	//action_type==2 then follower-1 offset=-1
	//offset=action_type*(-2)+3
	offset := action_type*(-2) + 3
	t := util.DbConn.Exec(
		`update 
			user 
		set 
			follow_count=follow_count+? where id=?`, offset, user_id)
	if t.Error != nil {
		util.DbConn.Rollback()
		util.DbConn.Unlock()
	}
	t = util.DbConn.Exec(
		`update 
			user 
		set 
			follower_count=follower_count+? where id=?`, offset, to_user_id)
	if t.Error != nil {
		util.DbConn.Rollback()
		util.DbConn.Unlock()
	}
	util.DbConn.Unlock()
}
