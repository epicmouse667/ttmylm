package mapper

import (
	"database/sql"
	"dou_sheng/pogo"
	"dou_sheng/util"
	"fmt"
)

func GetUserList(userList *map[string]int) {
	raw, _ := util.DbConn.DB().Query("select concat(name,password),id from user")
	var (
		key   string
		value int
	)
	for raw.Next() {
		raw.Scan(&key, &value)
		(*userList)[key] = value
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

func InsertUser(name string, password string) {
	unit, err := util.DbConn.DB().Prepare("insert into user(name,password,follow_count,follower_count)values(?,?,?,?)")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	_, err = unit.Exec(name, password, 0, 0)
	if err != nil {
		fmt.Println()
	}
}

func SelectUser(name string, password string) bool {

	rows := util.DbConn.DB().QueryRow("select name,password from user where name=? and password=?")

	//if err !=nil {
	//	fmt.Println("登录查询失败")
	//}
	//for rows.Next(){
	//	err=rows.Scan(&userName,&userPassword)
	//	if err!=nil {
	//		fmt.Println(&userName,&userPassword)
	//		fmt.Println(err.Error())
	//	}
	//}
	return rows.Err() != sql.ErrNoRows
}
