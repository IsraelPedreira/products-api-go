package routes

import (
	"marketplace-go/src/controllers"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(r *gin.RouterGroup) {
	r.POST("/", controllers.CreateProduct)
	r.GET("/:id", controllers.GetProductById)
	r.GET("/", controllers.GetAllProducts)
	r.PUT("/:id", controllers.UpdateProduct)
}
