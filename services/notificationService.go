package services

import (
	"WeatherApp/initializers"
	"WeatherApp/models"
	"WeatherApp/utils"
	"fmt"
	"time"
)

func ProcessSubscriptionsAndNotify() {
	var subscriptions []models.Subscription
	initializers.DB.Preload("Users").Find(&subscriptions)

	for _, sub := range subscriptions {
		if sub.LastSentAt != nil && sameDay(*sub.LastSentAt, time.Now()) {
			continue
		}

		var cache models.WeatherLog
		result := initializers.DB.First(&cache, "city = ?", sub.City)
		if result.Error != nil {
			fmt.Println("❌ Weather cache not found for:", sub.City)
			continue
		}
		if !utils.CheckCondition(cache.Temperature, sub.Condition) {
			continue
		}

		for _, user := range sub.Users {
			utils.SendEmail(user.Email, fmt.Sprintf("Weather alert for %s — %s (Current: %.1f°C)"))

			history := models.NotificationHistory{
				Email:     user.Email,
				City:      sub.City,
				Condition: sub.Condition,
				SentAt:    time.Now().Unix(),
			}
			initializers.DB.Create(&history)
		}

		now := time.Now()
		initializers.DB.Model(&sub).Update("LastSentAt", &now)
	}
}

func sameDay(a, b time.Time) bool {
	y1, m1, d1 := a.Date()
	y2, m2, d2 := b.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}
