package utils

import (
	"WeatherApp/initializers"
	"WeatherApp/models"
	"log"
)

func InitializeData() {
	var user models.User
	if err := initializers.DB.First(&user, "email = ?", "koltok.vitalik@gmail.com").Error; err != nil {
		user = models.User{Name: "Vitalii Koltok", Email: "koltok.vitalik@gmail.com"}
		initializers.DB.Create(&user)
		log.Println("Created tests user: koltok.vitalik@gmail.com")
	}

	sub := models.Subscription{City: "Warsaw", Condition: "temperature < 40"}
	initializers.DB.FirstOrCreate(&sub, models.Subscription{City: sub.City, Condition: sub.Condition})
	initializers.DB.Model(&user).Association("Subscriptions").Append(&sub)
	log.Println("Subscribed John to 'temperature < 40' in Warsaw")
}
