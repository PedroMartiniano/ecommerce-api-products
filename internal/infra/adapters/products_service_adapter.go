package adapters

import (
	"github.com/PedroMartiniano/ecommerce-api-products/internal/application/services"
	"github.com/PedroMartiniano/ecommerce-api-products/internal/configs"
	"github.com/PedroMartiniano/ecommerce-api-products/internal/infra/repositories"
)

func NewProductsServiceAdapter() *services.ProductsService {
	productsRepository := repositories.NewProductsRepository(configs.DB)
	stocksRepository := repositories.NewStocksRepository(configs.DB)
	redisRepository := repositories.NewRedisRepository(configs.Redis)
	productsService := services.NewProductsService(productsRepository, stocksRepository, redisRepository)

	return productsService
}
