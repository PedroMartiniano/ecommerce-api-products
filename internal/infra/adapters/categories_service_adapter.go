package adapters

import (
	"github.com/PedroMartiniano/ecommerce-api-products/internal/configs"
	"github.com/PedroMartiniano/ecommerce-api-products/internal/infra/repositories"
	"github.com/PedroMartiniano/ecommerce-api-products/internal/application/services"
)

func NewCategoriesServiceAdapter() *services.CategoriesService {
	categoriesRepository := repositories.NewCategoriesRepository(configs.DB)
	categoriesService := services.NewCategoriesService(categoriesRepository)

	return categoriesService
}
