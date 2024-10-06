package adapters

import (
	"github.com/PedroMartiniano/ecommerce-api-products/internal/repositories"
	"github.com/PedroMartiniano/ecommerce-api-products/internal/services"
)

func NewUserServiceAdapter() *services.ProductsService {
	productsRepository := repositories.NewProductsRepository()
	productsService := services.NewProductsService(productsRepository)

	return productsService
}
