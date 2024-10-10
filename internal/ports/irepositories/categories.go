package irepositories

import (
	"context"

	"github.com/PedroMartiniano/ecommerce-api-products/internal/models"
)

type ICategoriesRepository interface {
	Create(context.Context, models.Categories) (models.Categories, error)
	FindById(context.Context, string) (models.Categories, error)
	List(context.Context) ([]models.Categories, error)
	Update(context.Context, models.Products) (models.Products, error)
	Delete(context.Context, models.Products) (models.Products, error)
}
