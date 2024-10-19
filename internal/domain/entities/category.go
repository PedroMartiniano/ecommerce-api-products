package entities

import (
	"time"

	"github.com/PedroMartiniano/ecommerce-api-products/internal/domain/vo"
)

type Category struct {
	ID          vo.UUID   `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func NewCategory(id, name, description string, createdAt, updatedAt *time.Time) (Category, error) {
	if createdAt == nil {
		now := time.Now()
		createdAt = &now
	}

	if updatedAt == nil {
		now := time.Now()
		updatedAt = &now
	}

	return Category{
		ID:          vo.NewUUID(id),
		Name:        name,
		Description: description,
		CreatedAt:   *createdAt,
		UpdatedAt:   *updatedAt,
	}, nil
}

func (c *Category) GetID() string {
	return c.ID.GetValue()
}

func (c *Category) GetName() string {
	return c.Name
}

func (c *Category) GetDescription() string {
	return c.Description
}

func (c *Category) GetCreatedAt() time.Time {
	return c.CreatedAt
}

func (c *Category) GetUpdatedAt() time.Time {
	return c.UpdatedAt
}
