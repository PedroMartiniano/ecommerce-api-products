package routes

import (
	"github.com/PedroMartiniano/ecommerce-api-products/internal/adapters"
	"github.com/PedroMartiniano/ecommerce-api-products/internal/controllers"
	"github.com/gin-gonic/gin"
)

func productsRoutes(router *gin.RouterGroup) {
	productsService := adapters.NewProductsServiceAdapter()
	productsController := controllers.NewProductsController(productsService)

	router.POST("/", productsController.CreateProductHandler)
	router.GET("/", productsController.ListProductsHandler)
	router.GET("/:id", productsController.GetProductByIDHandler)
	router.PUT("/:id", productsController.UpdateProductHandler)
	router.DELETE("/:id", productsController.DeleteProductHandler)
}
