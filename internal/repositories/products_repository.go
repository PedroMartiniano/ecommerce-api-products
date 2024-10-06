package repositories

import (
	pr "github.com/PedroMartiniano/ecommerce-api-products/internal/ports/repositories"
)

type productsRepository struct{}

func NewProductsRepository() pr.IProductsRepository {
	return &productsRepository{}
}
