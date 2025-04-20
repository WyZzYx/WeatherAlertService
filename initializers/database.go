package initializers

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

var DB *gorm.DB

func ConnectToDb() {
	dsn := os.Getenv("DB_URL")
	var err error

	for i := 0; i < 10; i++ {
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			log.Println("Connected to database")
			return
		}
		log.Println("Failed DB connection, retrying...", err)
		time.Sleep(3 * time.Second)
	}

	log.Fatal("Could not connect to the database after 10 tries:", err)
}
