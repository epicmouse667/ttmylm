package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	initRouter(r)
	r.Run() // listen an
	// d serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
