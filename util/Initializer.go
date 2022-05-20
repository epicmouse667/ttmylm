package util

import (
	"github.com/jinzhu/gorm"
	"log"
	"os"
	"sync"
)
import _ "github.com/go-sql-driver/mysql"

var DbConn *gorm.DB
var Lock sync.Mutex

const (
	DBConfig  = "root:1234@tcp(127.0.0.1:3306)/dou_sheng?charset=utf8mb4&parseTime=true"
	SourceWeb = "http://192.168.199.211:8081"
)

func Init() {
	var err error
	DbConn, err = gorm.Open("mysql", DBConfig)
	if err != nil {
		log.Fatalln(err)
		os.Exit(-1)
	}
}
func Destu() {
	DbConn.Close()
}
