package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"terminal_commands/models"
)

var AppDb *gorm.DB

func InitializeDb() {

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&models.User{}, &models.Command{})

	if err != nil {
		panic(err)
	}

	AppDb = db
}
