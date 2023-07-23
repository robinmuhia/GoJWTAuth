package main

import (
	"github.com/gin-gonic/gin"
	"github.com/robinmuhia/GoJWTAuth/controllers"
	"github.com/robinmuhia/GoJWTAuth/initializers"
)


func init(){
	initializers.LoadEnvironmentVariables()
	initializers.ConnectToDatabase()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()
	r.POST("/signup", controllers.SignUp)
	r.Run() // listen and serve on 0.0.0.0:8080
}