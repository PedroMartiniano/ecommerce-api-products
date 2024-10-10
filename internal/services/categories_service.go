package services

import (
	"context"

	"github.com/PedroMartiniano/ecommerce-api-products/internal/models"
	pr "github.com/PedroMartiniano/ecommerce-api-products/internal/ports/irepositories"
)

type CategoriesService struct {
	categoriesRepository pr.ICategoriesRepository
}

func NewCategoriesService(categoriesRepository pr.ICategoriesRepository) *CategoriesService {
	return &CategoriesService{
		categoriesRepository: categoriesRepository,
	}
}

func (p *CategoriesService) CreateCategoryExecute(c context.Context, category models.Categories) (models.Categories, error) {
	newCategory, err := p.categoriesRepository.Create(c, category)

	return newCategory, err
}

func (p *CategoriesService) ListCategoriesExecute(c context.Context) ([]models.Categories, error) {
	categories, err := p.categoriesRepository.List(c)

	return categories, err
}
