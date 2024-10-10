package routes

import (
	"github.com/PedroMartiniano/ecommerce-api-products/internal/adapters"
	"github.com/PedroMartiniano/ecommerce-api-products/internal/controllers"
	"github.com/gin-gonic/gin"
)

func categoriesRoutes(router *gin.RouterGroup) {
	categoriesService := adapters.NewCategoriesServiceAdapter()
	categoriesController := controllers.NewCategoriesController(categoriesService)

	router.POST("/", categoriesController.CreateHandler)
	router.GET("/", categoriesController.ListHandler)
}
