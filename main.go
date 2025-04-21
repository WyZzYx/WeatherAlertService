package main

import (
	"WeatherApp/initializers"
	"WeatherApp/services"
	"fmt"
	"github.com/robfig/cron/v3"

	"WeatherApp/utils"
	"log"
	"os"
)

func init() {
	initializers.LoadEnvVariables()

	initializers.ConnectToDb()

	initializers.Migrate()

	utils.InitializeData()
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	r := services.SetupRouter()

	log.Printf("Server running on port %s", port)
	err := r.Run(":" + port)
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}

	c := cron.New()
	c.AddFunc("@daily", func() {
		fmt.Println("⏰ Running daily weather notifications...")
		services.ProcessSubscriptionsAndNotify()
	})
	c.Start()

}
