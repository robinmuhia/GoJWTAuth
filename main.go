package main

import (
	"github.com/gin-gonic/gin"
	"github.com/robinmuhia/GoJWTAuth/controllers"
	"github.com/robinmuhia/GoJWTAuth/initializers"
	"github.com/robinmuhia/GoJWTAuth/middleware"
)


func init(){
	initializers.LoadEnvironmentVariables()
	initializers.ConnectToDatabase()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()
	r.POST("/signup", controllers.SignUp)
	r.POST("/login",controllers.Login)
	r.GET("/logout",middleware.RequireAuth,controllers.Logout)
	r.GET("/home",middleware.RequireAuth,controllers.Home)
	r.Run() 
}