package services

import (
	"context"

	pr "github.com/PedroMartiniano/ecommerce-api-products/internal/application/ports/repositories"
	"github.com/PedroMartiniano/ecommerce-api-products/internal/domain/dto"
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

func (p *CategoriesService) CreateCategoryExecute(c context.Context, category dto.Category) (dto.Category, error) {
	categoryEntity, err := entities.NewCategory("", category.Name, category.Description, nil, nil)
	if err != nil {
		return dto.Category{}, err
	}
	newCategory, err := p.categoriesRepository.Create(c, categoryEntity)

	return dto.Category{
		ID:          newCategory.GetID(),
		Name:        newCategory.GetName(),
		Description: newCategory.GetDescription(),
		CreatedAt:   newCategory.GetCreatedAt(),
		UpdatedAt:   newCategory.GetUpdatedAt(),
	}, err
}

func (p *CategoriesService) ListCategoriesExecute(c context.Context) ([]dto.Category, error) {
	categories, err := p.categoriesRepository.List(c)

	categoriesDTO := []dto.Category{}
	for _, category := range categories {
		categoriesDTO = append(categoriesDTO, dto.Category{
			ID:          category.GetID(),
			Name:        category.GetName(),
			Description: category.GetDescription(),
			CreatedAt:   category.GetCreatedAt(),
			UpdatedAt:   category.GetUpdatedAt(),
		})
	}
	return categoriesDTO, err
}
