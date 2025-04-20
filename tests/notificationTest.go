package tests

import (
	"WeatherApp/models"
	"WeatherApp/utils"
	"testing"
	"time"
)

func TestNotificationConditionEvaluation(t *testing.T) {
	temp := 5.0
	cond := "temperature < 10"
	if !utils.CheckCondition(temp, cond) {
		t.Error("Expected condition to be true")
	}

	cond = "temperature > 10"
	if utils.CheckCondition(temp, cond) {
		t.Error("Expected condition to be false")
	}
}

func TestNotificationHistoryModel(t *testing.T) {
	history := models.NotificationHistory{
		Email:     "tests@example.com",
		City:      "Kyiv",
		Condition: "temperature < 10",
		SentAt:    time.Now().Unix(),
	}
	if history.Email == "" || history.City == "" || history.SentAt == 0 {
		t.Error("NotificationHistory model is missing required fields")
	}
}

func TestWeatherCacheModel(t *testing.T) {
	cache := models.WeatherLog{
		City:        "Kyiv",
		Temperature: 5.0,
		UpdatedAt:   time.Now().Unix(),
	}
	if cache.City == "" || cache.UpdatedAt == 0 {
		t.Error("WeatherCache model is missing required fields")
	}
}
