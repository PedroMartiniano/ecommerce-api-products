package ports

import (
	"context"

	"github.com/PedroMartiniano/ecommerce-api-products/internal/domain/entities"
)

type IProductsRepository interface {
	Create(context.Context, entities.Product) (entities.Product, error)
	FindById(context.Context, string) (entities.Product, error)
	List(context.Context) ([]entities.Product, error)
	Update(context.Context, entities.Product) (entities.Product, error)
	Delete(context.Context, string) error
}
