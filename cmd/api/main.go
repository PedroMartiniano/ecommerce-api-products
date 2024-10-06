package main

import (
	"fmt"
	"github.com/PedroMartiniano/ecommerce-api-products/internal/routes"

	"github.com/PedroMartiniano/ecommerce-api-products/internal/configs"
	"github.com/gin-gonic/gin"
)

func main() {
	configs.Init()

	gin.SetMode(gin.DebugMode)

	server := gin.Default()

	routes.InitRoutes(server)

	port := configs.GetEnv("PORT")
	server.Run(fmt.Sprintf(":%s", port))
}
