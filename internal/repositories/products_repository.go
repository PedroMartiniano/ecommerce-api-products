package repositories

import (
	"context"
	"database/sql"
	"errors"
	"github.com/PedroMartiniano/ecommerce-api-products/internal/configs"
	"github.com/PedroMartiniano/ecommerce-api-products/internal/models"
	pr "github.com/PedroMartiniano/ecommerce-api-products/internal/ports/irepositories"
	"github.com/google/uuid"
	"time"
)

type productsRepository struct {
	db *sql.DB
}

func NewProductsRepository(db *sql.DB) pr.IProductsRepository {
	return &productsRepository{
		db: db,
	}
}

func (p productsRepository) Create(ctx context.Context, product models.Products) (models.Products, error) {
	query := `INSERT INTO products(id, name, description, price, category_id, created_at, updated_at) VALUES($1, $2, $3, $4, $5, $6, $7);`

	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()
	id, _ := uuid.NewV7()
	product.ID = id.String()

	_, err := p.db.ExecContext(
		ctx,
		query,
		product.ID,
		product.Name,
		product.Description,
		product.Price,
		product.CategoryID,
		product.CreatedAt,
		product.UpdatedAt,
	)
	if err != nil {
		return models.Products{}, configs.NewError(configs.ErrInternalServer, err)
	}

	return product, nil
}

func (p productsRepository) FindById(ctx context.Context, id string) (models.Products, error) {
	query := `SELECT p.id, p.name, p.description, p.price, p.category_id, p.created_at, p.updated_at, s.quantity FROM products p INNER JOIN stocks s ON (p.id = s.product_id) WHERE p.id = $1;`

	row := p.db.QueryRowContext(ctx, query, id)

	product := models.Products{}

	err := row.Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Price,
		&product.CategoryID,
		&product.CreatedAt,
		&product.UpdatedAt,
		&product.Quantity,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Products{}, configs.NewError(configs.ErrNotFound, err)
		}
		return models.Products{}, configs.NewError(configs.ErrInternalServer, err)
	}

	return product, nil
}

func (p productsRepository) List(ctx context.Context) ([]models.Products, error) {
	query := `SELECT p.id, p.name, p.description, p.price, p.category_id, p.created_at, p.updated_at, s.quantity FROM products p INNER JOIN stocks s ON (p.id = s.product_id);`

	rows, err := p.db.QueryContext(ctx, query)
	if err != nil {
		return []models.Products{}, configs.NewError(configs.ErrInternalServer, err)
	}
	defer rows.Close()

	products := []models.Products{}

	for rows.Next() {
		product := models.Products{}
		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.CategoryID,
			&product.CreatedAt,
			&product.UpdatedAt,
			&product.Quantity,
		)
		if err != nil {
			return []models.Products{}, configs.NewError(configs.ErrInternalServer, err)
		}

		products = append(products, product)
	}

	return products, nil
}

func (p productsRepository) Update(ctx context.Context, product models.Products) (models.Products, error) {
	query := `UPDATE products SET name = $1, description = $2, price = $3, category_id = $4, updated_at = $5 WHERE id = $6;`

	product.UpdatedAt = time.Now()

	_, err := p.db.ExecContext(
		ctx,
		query,
		product.Name,
		product.Description,
		product.Price,
		product.CategoryID,
		product.UpdatedAt,
		product.ID,
	)
	if err != nil {
		return models.Products{}, configs.NewError(configs.ErrInternalServer, err)
	}

	return product, nil
}

func (p productsRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM products WHERE id = $1;`

	_, err := p.db.ExecContext(ctx, query, id)
	if err != nil {
		return configs.NewError(configs.ErrInternalServer, err)
	}

	return nil
}
