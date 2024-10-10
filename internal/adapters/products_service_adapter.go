package adapters

import (
	"github.com/PedroMartiniano/ecommerce-api-products/internal/configs"
	"github.com/PedroMartiniano/ecommerce-api-products/internal/repositories"
	"github.com/PedroMartiniano/ecommerce-api-products/internal/services"
)

func NewProductsServiceAdapter() *services.ProductsService {
	productsRepository := repositories.NewProductsRepository(configs.DB)
	stocksRepository := repositories.NewStocksRepository(configs.DB)
	productsService := services.NewProductsService(productsRepository, stocksRepository)

	return productsService
}
