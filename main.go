package main

import (
	"LibraryManagement/db"
	"LibraryManagement/router"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitRegist()
	gin := gin.Default()
	router.InitRouter(gin)
	gin.Run(":8080")
}
