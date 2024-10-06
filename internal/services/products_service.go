package services

import (
	"context"
	"github.com/PedroMartiniano/ecommerce-api-products/internal/models"
	"github.com/PedroMartiniano/ecommerce-api-products/internal/ports/repositories"
)

type ProductsService struct {
	productsRepository repositories.IProductsRepository
}

func NewProductsService(userRepository repositories.IProductsRepository) *ProductsService {
	return &ProductsService{
		productsRepository: userRepository,
	}
}

func (p *ProductsService) CreateExecute(c context.Context, product models.Products) (models.Products, error) {
	newProduct, err := p.productsRepository.Create(c, product)

	return newProduct, err
}
