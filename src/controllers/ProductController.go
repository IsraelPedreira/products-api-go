package controllers

import (
	"marketplace-go/src/database"
	"marketplace-go/src/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateProduct godoc
// @Summary Cria um novo produto
// @Description Adiciona um novo produto ao banco de dados
// @Tags produtos
// @Accept json
// @Produce json
// @Param product body models.Product true "Produto"
// @Success 201 {object} models.Product
// @Failure 400 {object} map[string]string
// @Router /products [post]
func CreateProduct(c *gin.Context) {
	var product models.Product

	if err := c.ShouldBindBodyWithJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&product)
	c.JSON(http.StatusCreated, product)

}

func GetAllProducts(c *gin.Context) {
	var products []models.Product

	if err := database.DB.Find(&products).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, products)
}

func GetProductById(c *gin.Context) {
	id := c.Param("id")

	var product models.Product

	if err := database.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, product)
}

func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product

	if err := database.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if err := c.ShouldBindBodyWithJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Save(&product)
	c.JSON(http.StatusCreated, product)

}
