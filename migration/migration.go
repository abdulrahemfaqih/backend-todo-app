package migration

import (
	"backend-mytodo/domain"
	"fmt"
	"gorm.io/gorm"
	"log"
)

func RunMigration(db *gorm.DB) {
	err := db.AutoMigrate(&domain.Todo{})

	if err != nil {
		log.Println(err)
	}

	fmt.Println("Database Migration Success")
}