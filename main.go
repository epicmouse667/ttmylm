package main

import (
	"github.com/gin-gonic/gin"
	"ttmylm/utils"
)

func main() {
	r := gin.Default()

	initRouter(r)
	utils.InitMysql()
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
