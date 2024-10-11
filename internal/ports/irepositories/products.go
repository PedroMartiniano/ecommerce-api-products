package irepositories

import (
	"context"

	"github.com/PedroMartiniano/ecommerce-api-products/internal/models"
)

type IProductsRepository interface {
	Create(context.Context, models.Product) (models.Product, error)
	FindById(context.Context, string) (models.Product, error)
	List(context.Context) ([]models.Product, error)
	Update(context.Context, models.Product) (models.Product, error)
	Delete(context.Context, string) error
}
