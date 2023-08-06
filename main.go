package main

import (
	"time"

	"github.com/gin-contrib/cors"
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
	r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:3000"},
        AllowMethods:     []string{"PUT", "PATCH", "GET", "DELETE"},
        AllowHeaders:     []string{"Origin"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        AllowOriginFunc: func(origin string) bool {
            return origin == "http://localhost:3000"
        },
        MaxAge: 12 * time.Hour,
    }))
	r.POST("/signup", controllers.SignUp)
	r.POST("/login",controllers.Login)
	r.GET("/logout",middleware.RequireAuth,controllers.Logout)
	r.GET("/home",middleware.RequireAuth,controllers.Home)
	r.Run() 
}