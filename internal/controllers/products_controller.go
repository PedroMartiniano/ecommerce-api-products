package controllers

import (
	"github.com/PedroMartiniano/ecommerce-api-products/internal/services"
	"github.com/gin-gonic/gin"
)

type ProductsController struct {
	productsService *services.ProductsService
}

func NewProductsController(productsService *services.ProductsService) *ProductsController {
	return &ProductsController{
		productsService: productsService,
	}
}

func (p *ProductsController) CreateHandler(c *gin.Context) {
}
