package main

import (
	"github.com/Mustafo770/blog-api/database" 
	"github.com/Mustafo770/blog-api/models"   
	"github.com/Mustafo770/blog-api/routes"   
	_ "gorm.io/driver/sqlite" )

func main() {
	
	database.Connect()
	database.DB.AutoMigrate(&models.Article{}, &models.Comment{}, &models.Like{})

	router := routes.SetupRouter()

	router.Run(":8080")
}
