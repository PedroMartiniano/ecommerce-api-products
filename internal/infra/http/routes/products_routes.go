package routes

import (
	"github.com/PedroMartiniano/ecommerce-api-products/internal/infra/adapters"
	"github.com/PedroMartiniano/ecommerce-api-products/internal/infra/http/controllers"
	"github.com/PedroMartiniano/ecommerce-api-products/internal/infra/http/middlewares"
	"github.com/gin-gonic/gin"
)

func productsRoutes(router *gin.RouterGroup) {
	productsService := adapters.NewProductsServiceAdapter()
	productsController := controllers.NewProductsController(productsService)

	router.POST("/", middlewares.VerifyToken, middlewares.VerifyStaff, productsController.CreateProductHandler)
	router.GET("/", productsController.ListProductsHandler)
	router.GET("/:id", productsController.GetProductByIDHandler)
	router.PUT("/:id", middlewares.VerifyToken, middlewares.VerifyStaff, productsController.UpdateProductHandler)
	router.DELETE("/:id", middlewares.VerifyToken, middlewares.VerifyStaff, productsController.DeleteProductHandler)
	router.PUT("/:id/stock", middlewares.VerifyToken, productsController.UpdateProductStockHandler)
	router.GET("/:id/stock", productsController.GetProductStockHandler)
}
