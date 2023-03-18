package database

import (
	"fmt"
	"log"

	"github.com/agusbasari29/xjx-biller-backend/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupDatabaseConnection() *gorm.DB {
	dbConfig := config.DbConfig()
	dsn := fmt.Sprintf("database/%s", dbConfig.File)
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to open database connection")
	}
	return db
}

func CloseDatabaseConnection(db *gorm.DB) {
	conn, err := db.DB()
	if err != nil {
		panic("Failed to close database connection")
	}
	conn.Close()
}
