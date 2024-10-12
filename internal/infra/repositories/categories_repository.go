package repositories

import (
	"context"
	"database/sql"
	"errors"
	"time"

	pr "github.com/PedroMartiniano/ecommerce-api-products/internal/application/ports"
	"github.com/PedroMartiniano/ecommerce-api-products/internal/configs"
	"github.com/PedroMartiniano/ecommerce-api-products/internal/domain/entities"
	"github.com/google/uuid"
)

type categoriesRepository struct {
	db *sql.DB
}

func NewCategoriesRepository(db *sql.DB) pr.ICategoriesRepository {
	return &categoriesRepository{
		db: db,
	}
}

func (c categoriesRepository) Create(ctx context.Context, category entities.Category) (entities.Category, error) {
	query := `INSERT INTO categories(id, name, description, created_at, updated_at) VALUES($1, $2, $3, $4, $5)`

	category.CreatedAt = time.Now()
	category.UpdatedAt = time.Now()
	uuid, _ := uuid.NewV7()
	category.ID = uuid.String()

	_, err := c.db.ExecContext(
		ctx,
		query,
		category.ID,
		category.Name,
		category.Description,
		category.CreatedAt,
		category.UpdatedAt,
	)
	if err != nil {
		return entities.Category{}, configs.NewError(configs.ErrInternalServer, err)
	}

	return category, nil
}

func (c categoriesRepository) FindById(ctx context.Context, id string) (entities.Category, error) {
	query := `SELECT id, name, description, created_at, updated_at FROM categories WHERE id = $1`

	row := c.db.QueryRowContext(
		ctx,
		query,
		id,
	)

	category := entities.Category{}

	err := row.Scan(
		&category.ID,
		&category.Name,
		&category.Description,
		&category.CreatedAt,
		&category.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entities.Category{}, configs.NewError(configs.ErrNotFound, err)
		}
		return entities.Category{}, configs.NewError(configs.ErrInternalServer, err)
	}

	return category, nil
}

func (c categoriesRepository) List(ctx context.Context) ([]entities.Category, error) {
	query := `SELECT id, name, description, created_at, updated_at FROM categories`

	rows, err := c.db.QueryContext(ctx, query)
	if err != nil {
		return []entities.Category{}, configs.NewError(configs.ErrInternalServer, err)
	}
	defer rows.Close()

	categories := []entities.Category{}
	for rows.Next() {
		category := entities.Category{}
		err := rows.Scan(
			&category.ID,
			&category.Name,
			&category.Description,
			&category.CreatedAt,
			&category.UpdatedAt,
		)
		if err != nil {
			return []entities.Category{}, configs.NewError(configs.ErrInternalServer, err)
		}

		categories = append(categories, category)
	}

	return categories, nil
}

func (c categoriesRepository) Update(ctx context.Context, products entities.Product) (entities.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (c categoriesRepository) Delete(ctx context.Context, products entities.Product) (entities.Product, error) {
	//TODO implement me
	panic("implement me")
}
