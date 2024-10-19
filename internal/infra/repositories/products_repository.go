package repositories

import (
	"context"
	"database/sql"
	"errors"
	"time"

	pr "github.com/PedroMartiniano/ecommerce-api-products/internal/application/ports"
	"github.com/PedroMartiniano/ecommerce-api-products/internal/configs"
	"github.com/PedroMartiniano/ecommerce-api-products/internal/domain/entities"
)

type productsRepository struct {
	db *sql.DB
}

func NewProductsRepository(db *sql.DB) pr.IProductsRepository {
	return &productsRepository{
		db: db,
	}
}

func (p productsRepository) Create(ctx context.Context, product entities.Product) (entities.Product, error) {
	query := `INSERT INTO products(id, name, description, price, category_id, created_at, updated_at) VALUES($1, $2, $3, $4, $5, $6, $7);`

	_, err := p.db.ExecContext(
		ctx,
		query,
		product.GetID(),
		product.GetName(),
		product.GetDescription(),
		product.GetPrice(),
		product.GetCategoryID(),
		product.GetCreatedAt(),
		product.GetUpdatedAt(),
	)
	if err != nil {
		return entities.Product{}, configs.NewError(configs.ErrInternalServer, err)
	}

	return product, nil
}

func (p productsRepository) FindById(ctx context.Context, id string) (entities.Product, error) {
	query := `SELECT p.id, p.name, p.description, p.price, p.category_id, p.created_at, p.updated_at, s.quantity FROM products p INNER JOIN stocks s ON (p.id = s.product_id) WHERE p.id = $1;`

	row := p.db.QueryRowContext(ctx, query, id)

	var productID, productName, productDescription, productCategoryID string
	var productCreatedAt, productUpdatedAt time.Time
	var productPrice float64
	var productQuantity int

	err := row.Scan(
		&productID,
		&productName,
		&productDescription,
		&productPrice,
		&productCategoryID,
		&productCreatedAt,
		&productUpdatedAt,
		&productQuantity,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entities.Product{}, configs.NewError(configs.ErrNotFound, err)
		}
		return entities.Product{}, configs.NewError(configs.ErrInternalServer, err)
	}

	product, err := entities.NewProduct(
		productID,
		productName,
		productDescription,
		productCategoryID,
		productQuantity,
		productPrice,
		&productCreatedAt,
		&productUpdatedAt,
	)
	if err != nil {
		return entities.Product{}, configs.NewError(configs.ErrInternalServer, err)
	}

	return product, nil
}

func (p productsRepository) List(ctx context.Context) ([]entities.Product, error) {
	query := `SELECT p.id, p.name, p.description, p.price, p.category_id, p.created_at, p.updated_at, s.quantity FROM products p INNER JOIN stocks s ON (p.id = s.product_id);`

	rows, err := p.db.QueryContext(ctx, query)
	if err != nil {
		return []entities.Product{}, configs.NewError(configs.ErrInternalServer, err)
	}
	defer rows.Close()

	products := []entities.Product{}

	for rows.Next() {
		var productID, productName, productDescription, productCategoryID string
		var productCreatedAt, productUpdatedAt time.Time
		var productPrice float64
		var productQuantity int
		err := rows.Scan(
			&productID,
			&productName,
			&productDescription,
			&productPrice,
			&productCategoryID,
			&productCreatedAt,
			&productUpdatedAt,
			&productQuantity,
		)
		if err != nil {
			return []entities.Product{}, configs.NewError(configs.ErrInternalServer, err)
		}

		product, err := entities.NewProduct(
			productID,
			productName,
			productDescription,
			productCategoryID,
			productQuantity,
			productPrice,
			&productCreatedAt,
			&productUpdatedAt,
		)
		if err != nil {
			return []entities.Product{}, configs.NewError(configs.ErrInternalServer, err)
		}

		products = append(products, product)
	}

	return products, nil
}

func (p productsRepository) Update(ctx context.Context, product entities.Product) (entities.Product, error) {
	query := `UPDATE products SET name = $1, description = $2, price = $3, category_id = $4, updated_at = $5 WHERE id = $6;`

	product.UpdatedAt = time.Now()

	_, err := p.db.ExecContext(
		ctx,
		query,
		product.GetName(),
		product.GetDescription(),
		product.GetPrice(),
		product.GetCategoryID(),
		product.GetUpdatedAt(),
		product.GetID(),
	)
	if err != nil {
		return entities.Product{}, configs.NewError(configs.ErrInternalServer, err)
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
