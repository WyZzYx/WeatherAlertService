package controllers

import (
	"WeatherApp/initializers"
	"WeatherApp/models"
	"WeatherApp/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type SubscriptionInput struct {
	Email     string `json:"email" binding:"required,email"`
	City      string `json:"city" binding:"required"`
	Condition string `json:"condition" binding:"required"`
}

func SubscribeUser(c *gin.Context) {
	var input SubscriptionInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	result := initializers.DB.Where("email = ?", input.Email).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	sub := models.Subscription{
		City:      input.City,
		Condition: input.Condition,
	}
	initializers.DB.FirstOrCreate(&sub, models.Subscription{City: input.City, Condition: input.Condition})
	initializers.DB.Model(&user).Association("Subscriptions").Append(&sub)
	initializers.DB.Create(&models.NotificationHistory{
		Email:     user.Email,
		City:      sub.City,
		Condition: sub.Condition,
		SentAt:    time.Now().Unix(),
	})
	utils.SendEmail(user.Email, fmt.Sprintf("Subscribed to weather alerts in %s with condition: %s", sub.City, sub.Condition))
	c.JSON(http.StatusOK, gin.H{"message": "Subscription created"})
}
