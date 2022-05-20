package util

import (
	"github.com/jinzhu/gorm"
	"log"
	"os"
)
import _ "github.com/go-sql-driver/mysql"

var DbConn *gorm.DB

const (
	DBConfig  = "root:1234@tcp(192.168.199.183:3306)/dou_sheng?charset=utf8mb4&parseTime=true"
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
