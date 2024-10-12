package adapters

import (
	"github.com/PedroMartiniano/ecommerce-api-products/internal/configs"
	"github.com/PedroMartiniano/ecommerce-api-products/internal/infra/repositories"
	"github.com/PedroMartiniano/ecommerce-api-products/internal/application/services"
)

func NewProductsServiceAdapter() *services.ProductsService {
	productsRepository := repositories.NewProductsRepository(configs.DB)
	stocksRepository := repositories.NewStocksRepository(configs.DB)
	productsService := services.NewProductsService(productsRepository, stocksRepository)

	return productsService
}
