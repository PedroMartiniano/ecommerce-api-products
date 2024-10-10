package irepositories

import (
	"context"

	"github.com/PedroMartiniano/ecommerce-api-products/internal/models"
)

type IProductsRepository interface {
	Create(context.Context, models.Products) (models.Products, error)
	FindById(context.Context, string) (models.Products, error)
	List(context.Context) ([]models.Products, error)
	Update(context.Context, models.Products) (models.Products, error)
	Delete(context.Context, string) error
}
