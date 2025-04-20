package controllers

import (
	"WeatherApp/initializers"
	"WeatherApp/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserInput struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}

func CreateUser(c *gin.Context) {
	var input UserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := models.User{Name: input.Name, Email: input.Email}
	initializers.DB.Create(&user)
	c.JSON(http.StatusCreated, user)
}
