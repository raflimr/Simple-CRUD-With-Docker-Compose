package main

import (
	"d-crud/config"
	"d-crud/routes"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.ConnectDB()
)

func main() {
	defer config.CloseDatabaseConnection(db)
	router := gin.Default()

	routes.RouterUser(router)

	router.Run()
}
