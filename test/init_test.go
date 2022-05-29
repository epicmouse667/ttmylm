package test

import (
	"database/sql"
	"dou_sheng/mapper"
	"dou_sheng/pogo"
	"dou_sheng/util"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"testing"
	"time"
)

type FeedResponse struct {
	pogo.Response
	VideoList []pogo.Video `json:"video_list,omitempty"`
	NextTime  int64        `json:"next_time,omitempty"`
}

func TestRun(t *testing.T) {
	r := gin.Default()
	util.Init()
	r.Static("/static", "./public")

	apiRouter := r.Group("/douyin")

	// basic apis
	apiRouter.GET("/feed/", func(c *gin.Context) {
		list := *mapper.GetFeedList(1)
		for i := 0; i < len(list); i++ {
			id := list[i].AuthorID
			list[i].Author = mapper.GetUserByID(id)
		}
		fmt.Println(list)
		c.JSON(http.StatusOK, FeedResponse{
			Response:  pogo.Response{StatusCode: 0},
			VideoList: list,
			NextTime:  time.Now().Unix(),
		})
	})

	r.Run()
}

func TestDB(t *testing.T) {
	util.Init()
	stmt, err := util.DbConn.DB().Prepare("select * from user_favorite where user_id=? and video_id=?")
	if err != nil {
		fmt.Println("youwenti")
	}
	var res = false
	err1 := stmt.QueryRow(3, 3).Scan(&res)
	if err1 == sql.ErrNoRows {
		fmt.Println("no result")
	} else {
		fmt.Println("has result")
	}

	//x:=stmt.QueryRow(5,2)
	//if(x!=nil){
	//	println("not null")
	//} else {
	//	println("null")
	//}
}
func TestDo(t *testing.T) {
	util.Init()
	fmt.Println(*mapper.GetFavoriteList(1, 1))
}
