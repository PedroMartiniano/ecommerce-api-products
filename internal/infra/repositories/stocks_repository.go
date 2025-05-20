package repositories

import (
	"context"
	"database/sql"
	"time"

	pr "github.com/PedroMartiniano/ecommerce-api-products/internal/application/ports/repositories"
	"github.com/PedroMartiniano/ecommerce-api-products/internal/configs"
	"github.com/PedroMartiniano/ecommerce-api-products/internal/domain/entities"
)

type StocksRepository struct {
	db *sql.DB
}

func NewStocksRepository(db *sql.DB) pr.IStocksRepository {
	return &StocksRepository{
		db: db,
	}
}

func (s *StocksRepository) Create(ctx context.Context, stock entities.Stock) (entities.Stock, error) {
	query := `INSERT INTO stocks (id, product_id, quantity, updated_at) VALUES ($1, $2, $3, $4);`

	_, err := s.db.ExecContext(
		ctx,
		query,
		stock.GetID(),
		stock.GetProductID(),
		stock.GetQuantity(),
		stock.GetUpdatedAt(),
	)
	if err != nil {
		return entities.Stock{}, configs.NewError(configs.ErrInternalServer, err)
	}

	return stock, nil
}

func (s *StocksRepository) Update(ctx context.Context, stock entities.Stock) (entities.Stock, error) {
	query := `UPDATE stocks SET quantity = $1, updated_at = $2 WHERE product_id = $3;`

	_, err := s.db.ExecContext(
		ctx,
		query,
		stock.GetQuantity(),
		stock.GetUpdatedAt(),
		stock.GetProductID(),
	)
	if err != nil {
		return entities.Stock{}, configs.NewError(configs.ErrInternalServer, err)
	}

	return stock, nil
}

func (s *StocksRepository) DeleteByProductID(c context.Context, productID string) error {
	query := `DELETE FROM stocks WHERE product_id = $1;`

	_, err := s.db.ExecContext(c, query, productID)
	if err != nil {
		return configs.NewError(configs.ErrInternalServer, err)
	}

	return nil
}

func (s *StocksRepository) GetByProductID(c context.Context, productID string) (entities.Stock, error) {
	query := `SELECT id, product_id, quantity, updated_at FROM stocks WHERE product_id = $1;`

	var stock entities.Stock
	row := s.db.QueryRowContext(c, query, productID)

	var id, stockProductID string
	var quantity int
	var updatedAt time.Time

	err := row.Scan(
		&id,
		&stockProductID,
		&quantity,
		&updatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return entities.Stock{}, configs.NewError(configs.ErrNotFound, err)
		}
		return entities.Stock{}, configs.NewError(configs.ErrInternalServer, err)
	}

	stock, err = entities.NewStock(id, stockProductID, quantity, &updatedAt)
	if err != nil {
		return entities.Stock{}, configs.NewError(configs.ErrBadRequest, err)
	}

	return stock, nil
}
