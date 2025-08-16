package controllers

import (
	"marketplace-go/src/configuration/database"
	"marketplace-go/src/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Cria um novo produto
// @Description Adiciona um novo produto ao banco de dados
// @Tags produtos
// @Accept json
// @Produce json
// @Param product body models.Product true "Produto"
// @Success 201 {object} models.Product
// @Failure 400
// @Router /product [post]
func CreateProduct(c *gin.Context) {
	var product models.Product

	if err := c.ShouldBindBodyWithJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&product)
	c.JSON(http.StatusCreated, product)

}

// @Summary Lista todos os produtos
// @Description Retorna todos os produtos cadastrados
// @Tags produtos
// @Produce json
// @Success 200 {array} models.Product
// @Failure 500
// @Router /product [get]
func GetAllProducts(c *gin.Context) {
	var products []models.Product

	if err := database.DB.Find(&products).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, products)
}

// @Summary Busca produto por ID
// @Description Retorna um produto específico pelo ID
// @Tags produtos
// @Produce json
// @Param id path int true "ID do produto"
// @Success 200 {object} models.Product
// @Failure 404
// @Router /product/{id} [get]
func GetProductById(c *gin.Context) {
	id := c.Param("id")

	var product models.Product

	if err := database.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, product)
}

// @Summary Atualiza um produto
// @Description Atualiza os dados de um produto existente
// @Tags produtos
// @Accept json
// @Produce json
// @Param id path int true "ID do produto"
// @Param product body models.Product true "Dados do produto"
// @Success 200 {object} models.Product
// @Failure 400
// @Failure 404
// @Router /product/{id} [put]
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

// @Summary Deleta um produto
// @Description Remove um produto do banco de dados
// @Tags produtos
// @Param id path int true "ID do produto"
// @Success 200
// @Failure 404
// @Router /product/{id} [delete]
func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product

	if err := database.DB.Delete(&product, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, "Deleted successfully")
}
