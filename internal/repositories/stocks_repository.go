package repositories

import (
	"context"
	"database/sql"
	"time"

	"github.com/PedroMartiniano/ecommerce-api-products/internal/configs"
	"github.com/PedroMartiniano/ecommerce-api-products/internal/models"
	pr "github.com/PedroMartiniano/ecommerce-api-products/internal/ports/irepositories"
	"github.com/google/uuid"
)

type StocksRepository struct {
	db *sql.DB
}

func NewStocksRepository(db *sql.DB) pr.IStocksRepository {
	return &StocksRepository{
		db: db,
	}
}

func (s *StocksRepository) Create(ctx context.Context, stock models.Stocks) (models.Stocks, error) {
	query := `INSERT INTO stocks (id, product_id, quantity, updated_at) VALUES ($1, $2, $3, $4);`

	stock.UpdatedAt = time.Now()
	id, _ := uuid.NewV7()
	stock.ID = id.String()

	_, err := s.db.ExecContext(
		ctx,
		query,
		stock.ID,
		stock.ProductID,
		stock.Quantity,
		stock.UpdatedAt,
	)
	if err != nil {
		return models.Stocks{}, configs.NewError(configs.ErrInternalServer, err)
	}

	return stock, nil
}

func (s *StocksRepository) Update(ctx context.Context, stock models.Stocks) (models.Stocks, error) {
	query := `UPDATE stocks SET quantity = $1, updated_at = $2 WHERE product_id = $3;`

	stock.UpdatedAt = time.Now()

	_, err := s.db.ExecContext(
		ctx,
		query,
		stock.Quantity,
		stock.UpdatedAt,
		stock.ProductID,
	)
	if err != nil {
		return models.Stocks{}, configs.NewError(configs.ErrInternalServer, err)
	}

	return stock, nil
}
