package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/robinmuhia/GoJWTAuth/models"
)

func Home(c *gin.Context) {
	user, _ := c.Get("user")
	
	c.JSON(http.StatusOK,gin.H{
		"id": user.(models.User).ID,
	})
}