package main

import (
	"fmt"
	"github.com/PedroMartiniano/ecommerce-api-products/internal/infra/http/routes"

	"github.com/PedroMartiniano/ecommerce-api-products/internal/configs"
	"github.com/gin-gonic/gin"
)

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	configs.Init()

	gin.SetMode(gin.DebugMode)

	server := gin.Default()

	routes.InitRoutes(server)

	port := configs.GetEnv("PORT")
	server.Run(fmt.Sprintf(":%s", port))
}
