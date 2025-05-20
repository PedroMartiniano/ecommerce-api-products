package repositories

import (
	"context"

	"github.com/PedroMartiniano/ecommerce-api-products/internal/domain/entities"
)

type ICategoriesRepository interface {
	Create(context.Context, entities.Category) (entities.Category, error)
	FindById(context.Context, string) (entities.Category, error)
	List(context.Context) ([]entities.Category, error)
	Update(context.Context, entities.Product) (entities.Product, error)
	Delete(context.Context, entities.Product) (entities.Product, error)
}
