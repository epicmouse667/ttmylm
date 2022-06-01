package mapper

//用来处理评论相关操作
import (
	"dou_sheng/pogo"
	"dou_sheng/util"
	"time"
)

//获得所有评论列表
func GetCommentList(video_id int) *[]pogo.Comment {
	//commentlist的容量上限为1000
	commentlist := make([]pogo.Comment, 0, 1000)
	//通过video_id获得该视频的所有评论
	raw, _ := util.DbConn.DB().Query(
		`select 
			comment.id,user.id,content,create_date,name,follow_count,follower_count
		from 
			comment
			join user on comment.user_id=user.id
		 where 
		 	video_id=?`, video_id)
	for raw.Next() {
		var comment pogo.Comment
		raw.Scan(&comment.Id, &comment.User.Id, &comment.Content,
			comment.CreateDate, comment.User.Name,
			comment.User.FollowCount, comment.User.FollowerCount)
		comment.User.IsFollow = GetUserRelation(comment.User.Id,
			GetAuthorIdByVideoId(video_id))
		commentlist = append(commentlist, comment)
	}
	return &commentlist
}

//评论的操作 此时action_type==1
func AddComment(user_id int, video_id int, comment_text string) *pogo.Comment {
	util.DbConn.Lock()
	t := util.DbConn.Exec(
		`insert into 
			comment(user_id,video_id,content,create_date)
		values(
			?,?,?,?
		)`, user_id, video_id, comment_text, time.Now().Format("2006-01-02 15:04:05"))
	if t.Error != nil {
		util.DbConn.Rollback()
		util.DbConn.Unlock()
		return nil
	}
	util.DbConn.Unlock()
	//这边的返回信息需要comment_id(本身在表中是自增的)
	//所以又用sql语句搜索了一遍，感觉效率有点低，可以改进
	var comment pogo.Comment
	raw, _ := util.DbConn.DB().Query(
		`select 
		comment.id,user.id,content,create_date,name,follow_count,follower_count
	from 
		comment
		join user on comment.user_id=user.id
	 where 
		 content=?`, comment_text)
	for raw.Next() {
		raw.Scan(&comment.Id, &comment.User.Id, &comment.Content,
			&comment.CreateDate, &comment.User.Name,
			&comment.User.FollowCount, &comment.User.FollowerCount)
		comment.User.IsFollow = GetUserRelation(comment.User.Id,
			GetAuthorIdByVideoId(video_id))
	}
	return &comment
}

//删除评论 此时action_type==2
func DeleteComment(user_id int, video_id int, comment_id int) *pogo.Comment {
	util.DbConn.Lock()
	comment := pogo.Comment{}
	raw, _ := util.DbConn.DB().Query(
		`select 
		comment.id,user.id,content,create_date,name,follow_count,follower_count
	from 
		comment
		join user on comment.user_id=user.id
	 where 
		 comment.id=?`, comment_id)
	for raw.Next() {
		raw.Scan(&comment.Id, &comment.User.Id, &comment.Content,
			&comment.CreateDate, &comment.User.Name,
			&comment.User.FollowCount, &comment.User.FollowerCount)
		comment.User.IsFollow = GetUserRelation(comment.User.Id,
			GetAuthorIdByVideoId(video_id))
	}
	t := util.DbConn.Exec(
		`delete from
			comment
		where
			id=?
		`, comment_id)
	if t.Error != nil {
		util.DbConn.Rollback()
		util.DbConn.Unlock()
	}
	util.DbConn.Unlock()
	return &comment
}
