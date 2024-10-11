package irepositories

import (
	"context"

	"github.com/PedroMartiniano/ecommerce-api-products/internal/models"
)

type ICategoriesRepository interface {
	Create(context.Context, models.Category) (models.Category, error)
	FindById(context.Context, string) (models.Category, error)
	List(context.Context) ([]models.Category, error)
	Update(context.Context, models.Product) (models.Product, error)
	Delete(context.Context, models.Product) (models.Product, error)
}
