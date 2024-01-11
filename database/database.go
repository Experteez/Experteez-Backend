package database

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	var err error
	dsn := "postgresql://adminexperteez:h720OBc_xRyUIF4y-DAB1A@experteez-development-8100.8nk.cockroachlabs.cloud:26257/defaultdb?sslmode=verify-full"

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database", err)
	}

	var now time.Time
	DB.Raw("SELECT NOW()").Scan(&now)

	fmt.Println(now)
	fmt.Println("Connection Opened to Database")
}
