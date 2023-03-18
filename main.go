package main

import (
	"os"

	"github.com/agusbasari29/xjx-biller-backend/database"
	"github.com/agusbasari29/xjx-biller-backend/entity"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = database.SetupDatabaseConnection()
)

func main() {
	defer database.CloseDatabaseConnection(db)
	db.AutoMigrate(&entity.Users{}, &entity.Clients{}, &entity.Products{})
	g := gin.Default()
	g.Run(os.Getenv("SERVER_PORT"))
}
