package initializers

import (
	"WeatherApp/controllers"
	"WeatherApp/models"
)

func Migrate() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Subscription{})
	DB.AutoMigrate(&models.NotificationHistory{})
	DB.AutoMigrate(&controllers.WeatherResponse{})
	DB.AutoMigrate(&models.WeatherLog{})
}
