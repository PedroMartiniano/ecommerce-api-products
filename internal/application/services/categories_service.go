package services

import (
	"context"

	pr "github.com/PedroMartiniano/ecommerce-api-products/internal/application/ports"
	"github.com/PedroMartiniano/ecommerce-api-products/internal/domain/entities"
)

type CategoriesService struct {
	categoriesRepository pr.ICategoriesRepository
}

func NewCategoriesService(categoriesRepository pr.ICategoriesRepository) *CategoriesService {
	return &CategoriesService{
		categoriesRepository: categoriesRepository,
	}
}

func (p *CategoriesService) CreateCategoryExecute(c context.Context, category entities.Category) (entities.Category, error) {
	newCategory, err := p.categoriesRepository.Create(c, category)

	return newCategory, err
}

func (p *CategoriesService) ListCategoriesExecute(c context.Context) ([]entities.Category, error) {
	categories, err := p.categoriesRepository.List(c)

	return categories, err
}
