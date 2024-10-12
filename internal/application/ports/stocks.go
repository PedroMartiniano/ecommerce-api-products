package ports

import (
	"context"

	"github.com/PedroMartiniano/ecommerce-api-products/internal/domain/entities"
)

type IStocksRepository interface {
	Create(context.Context, entities.Stock) (entities.Stock, error)
	Update(context.Context, entities.Stock) (entities.Stock, error)
	DeleteByProductID(context.Context, string) error
	GetByProductID(context.Context, string) (entities.Stock, error)
}
