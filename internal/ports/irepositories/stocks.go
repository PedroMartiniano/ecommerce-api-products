package irepositories

import (
	"context"

	"github.com/PedroMartiniano/ecommerce-api-products/internal/models"
)

type IStocksRepository interface {
	Create(context.Context, models.Stock) (models.Stock, error)
	Update(context.Context, models.Stock) (models.Stock, error)
}
