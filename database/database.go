package database

import (
	"fmt"

	"github.com/100lvlmaster/5-chan-go/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	// DBConn is a database connection singleton
	DBConn *gorm.DB
)

func InitDb() *gorm.DB {

	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		panic("Could not initialize database")
	}
	migrate(db)
	return db
}

func migrate(db *gorm.DB) {
	err := db.AutoMigrate(&models.Post{}, &models.Reply{})
	if err != nil {
		fmt.Print("Could not init database")
		return
	}
	fmt.Print("Database migrated")
}
