package main

import (
	"log"
	docs "marketplace-go/docs"
	"marketplace-go/src/configuration/database"
	"marketplace-go/src/models"
	"marketplace-go/src/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title API de Produtos
// @version 1.0
// @description Esta é a documentação da API de Produtos
// @host      localhost:8080
// @BasePath /
func main() {

	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	database.Connect()
	database.DB.AutoMigrate(&models.Product{})

	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	docs.SwaggerInfo.BasePath = "/"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	productRoutes := router.Group("/product")
	routes.ProductRoutes(productRoutes)

	router.Run(":8080")
}
