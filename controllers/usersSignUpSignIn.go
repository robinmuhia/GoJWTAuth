package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/robinmuhia/GoJWTAuth/initializers"
	"github.com/robinmuhia/GoJWTAuth/models"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context){
	var body struct{
		Email string
		Password string
		Name string
		Image string
	}
	
	if err := c.Bind(&body); err != nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"error":"Failed to get all values from body",
		})
		return
	}

	hashedPassword,err := bcrypt.GenerateFromPassword([]byte(body.Password),10)

	if err != nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"error":"Failed to hash password",
		})
		return
	}
	user := models.User{
		Name: body.Name,
		Password: string(hashedPassword),
		Email: body.Email,
		Image: body.Image,
	}
	result := initializers.DB.Create(&user)
	if result.Error != nil{
		c.JSON(http.StatusMethodNotAllowed,gin.H{
			"error":"Failed to create user",
		})
		return
	}
	c.JSON(http.StatusCreated,gin.H{})
}