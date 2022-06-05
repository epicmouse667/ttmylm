package util

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"

	_ "github.com/go-sql-driver/mysql"
)

var DbConn *gorm.DB

const (

	//DBConfig = "root:a@tcp(127.0.0.1:3306)/dou_sheng?charset=utf8mb4&parseTime=true"
	DBConfig  = "ds:1234@tcp(139.224.105.6:3306)/dou_sheng?charset=utf8mb4&parseTime=true"
	SourceWeb = "http://192.168.199.183:8081"
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
