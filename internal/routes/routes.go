package routes

import "github.com/gin-gonic/gin"

func InitRoutes(server *gin.Engine) {
	productsRouter := server.Group("/products")
	productsRoutes(productsRouter)

	categoriesRouter := server.Group("/categories")
	categoriesRoutes(categoriesRouter)
}
