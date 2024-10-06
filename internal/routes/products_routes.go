package routes

import (
	"github.com/PedroMartiniano/ecommerce-api-products/internal/adapters"
	"github.com/PedroMartiniano/ecommerce-api-products/internal/controllers"
	"github.com/gin-gonic/gin"
)

func productsRoutes(router *gin.RouterGroup) {
	productsService := adapters.NewUserServiceAdapter()
	productsController := controllers.NewProductsController(productsService)

	router.POST("/", productsController.CreateHandler)
}
