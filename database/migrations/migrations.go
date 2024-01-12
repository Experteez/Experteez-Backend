package migrations

import (
	"Experteez-Backend/database"
	"Experteez-Backend/model/entity"
	"fmt"

	"log"
)

func RunMigrations() {
	if database.DB == nil {
		fmt.Printf("Database connection: %v\n", database.DB)
		log.Fatal("Database connection is nil")
	}

	err := database.DB.AutoMigrate(&entity.User{}, &entity.Talent{}, &entity.Mentor{}, &entity.Partner{}, &entity.Admin{}, &entity.Project{})

	if err != nil {
		log.Fatal("Failed to migrate database", err)
	}

	fmt.Println("Migration run successfully")
}
