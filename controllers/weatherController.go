package controllers

import (
	"WeatherApp/initializers"
	"WeatherApp/models"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type WeatherResponse struct {
	Location struct {
		Name string `json:"name"`
	} `json:"location"`
	Current struct {
		TempC float64 `json:"temp_c"`
	} `json:"current"`
}

func GetWeather(c *gin.Context) {
	city := c.Query("city")
	if city == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing ?city parameter"})
		return
	}

	apiKey := os.Getenv("WEATHER_API_KEY")
	if apiKey == "" {
		log.Println("WEATHER_API_KEY is not set")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Missing API key"})
		return
	}

	url := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s", apiKey, city)
	log.Println("Fetching URL:", url)

	resp, err := http.Get(url)
	if err != nil {
		log.Println(" HTTP request failed:", err)
		c.JSON(http.StatusBadGateway, gin.H{"error": "Request to weather API failed"})
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	log.Println("Raw response:", string(body))

	if resp.StatusCode != http.StatusOK {
		c.JSON(resp.StatusCode, gin.H{"error": "Failed to fetch weather data"})
		return
	}

	var weather WeatherResponse
	if err := json.Unmarshal(body, &weather); err != nil {
		log.Println("Failed to parse response:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Parsing error"})
		return
	}
	initializers.DB.Save(&models.WeatherLog{
		City:        weather.Location.Name,
		Temperature: weather.Current.TempC,
		UpdatedAt:   time.Now().Unix(),
	})

	c.JSON(http.StatusOK, gin.H{
		"city":        weather.Location.Name,
		"temperature": fmt.Sprintf("%.1f°C", weather.Current.TempC),
	})
}
