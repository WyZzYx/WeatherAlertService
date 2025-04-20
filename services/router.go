package services

import (
	"WeatherApp/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/weather", controllers.GetWeather)
	r.POST("/users", controllers.CreateUser)
	r.POST("/subscribe", controllers.SubscribeUser)

	return r
}
