package irepositories

import (
	"context"

	"github.com/PedroMartiniano/ecommerce-api-products/internal/models"
)

type IStocksRepository interface {
	Create(context.Context, models.Stocks) (models.Stocks, error)
	Update(context.Context, models.Stocks) (models.Stocks, error)
}
