package router

import (
	"LibraryManagement/api"
	"LibraryManagement/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter(gin *gin.Engine) {
	gin.Use(middleware.Cors())
	login :=gin.Group("/login")
	{
		login.POST("/sendEmail",api.EmailApi.SendEmail)
		login.POST("/register",api.User.Register)
		login.GET("/user",api.User.Login)
	}
	user :=gin.Group("/user")
	{
		user.GET("/list",api.User.List)
		user.POST("/update",api.User.Update)
		user.GET("/searchEmail",api.User.SearchEmail)
	}
}
