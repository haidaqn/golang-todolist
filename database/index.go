package database

import (
	"log"
	"todolist/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase(){
	var err error
	DB, err = gorm.Open(sqlite.Open("todolist.db"), &gorm.Config{})

	if err != nil {
		log.Panic("failed to connect database")
	}

	DB.AutoMigrate(&models.Todo{})
}