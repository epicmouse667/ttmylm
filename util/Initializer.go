package util

import (
	"database/sql"
	"github.com/jinzhu/gorm"
	"log"
	"os"
	"sync"
)
import _ "github.com/go-sql-driver/mysql"

var DbConn *gorm.DB
var Lock sync.Mutex
var Stmt *sql.Stmt

const (
	DBConfig  = "root:1234@tcp(127.0.0.1:3306)/dou_sheng?charset=utf8mb4&parseTime=true"
	SourceWeb = "http://192.168.199.183:8081"
)

func InitSQL() {
	var err error
	Stmt, err = DbConn.DB().Prepare("select user_id from user_favorite where user_id=? and video_id=?")
	if err != nil {
		log.Fatalln(err)
		return
	}
}
func Init() {
	var err error
	DbConn, err = gorm.Open("mysql", DBConfig)
	if err != nil {
		log.Fatalln(err)
		os.Exit(-1)
	}
	InitSQL()
}
func Destu() {
	DbConn.Close()
}
