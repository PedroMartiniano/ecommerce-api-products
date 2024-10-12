package routes

import (
	"github.com/PedroMartiniano/ecommerce-api-products/internal/configs"
	"github.com/PedroMartiniano/ecommerce-api-products/internal/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRoutes(server *gin.Engine) {
	configs.SwaggerConfigure(docs.SwaggerInfo)

	productsRouter := server.Group("/products")
	productsRoutes(productsRouter)

	categoriesRouter := server.Group("/categories")
	categoriesRoutes(categoriesRouter)

	server.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
