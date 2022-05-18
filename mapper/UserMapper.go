package mapper

import (
	"dou_sheng/pogo"
	"dou_sheng/util"
	"fmt"
)

func GetUserList() []pogo.User {
	var userList []pogo.User
	err := util.DbConn.Raw("select id,name,follow_count,follower_count from user").Scan(&userList)
	if err != nil {
		fmt.Println("err in sql")
		return []pogo.User{}
	}
	fmt.Println("from db: ", userList)
	return userList
}
func GetUserByID(idd int) pogo.User {
	var user pogo.User
	util.DbConn.Raw("select id,name,follow_count,follower_count from user where id= ? ", idd).Scan(&user)
	//if(err!=nil){
	//	fmt.Println("ERR in querybyid")
	//	return pogo.User{}
	//}
	return user
}
