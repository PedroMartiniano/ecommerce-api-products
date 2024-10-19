package entities

import (
	"time"

	"github.com/PedroMartiniano/ecommerce-api-products/internal/domain/vo"
)

type Product struct {
	ID          vo.UUID     `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	CategoryID  vo.UUID     `json:"category_id"`
	Quantity    vo.Quantity `json:"quantity"`
	Price       vo.Price    `json:"price"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
}

func NewProduct(id, name, description, categoryID string, quantity int, price float64, createdAt, updatedAt *time.Time) (Product, error) {
	if createdAt == nil {
		now := time.Now()
		createdAt = &now
	}

	if updatedAt == nil {
		now := time.Now()
		updatedAt = &now
	}

	quantityVO, err := vo.NewQuantity(quantity)
	if err != nil {
		return Product{}, err
	}

	priceVO, err := vo.NewPrice(price)
	if err != nil {
		return Product{}, err
	}

	return Product{
		ID:          vo.NewUUID(id),
		Name:        name,
		Description: description,
		CategoryID:  vo.NewUUID(categoryID),
		Quantity:    quantityVO,
		Price:       priceVO,
		CreatedAt:   *createdAt,
		UpdatedAt:   *updatedAt,
	}, nil
}

func (p *Product) GetID() string {
	return p.ID.GetValue()
}

func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) GetDescription() string {
	return p.Description
}

func (p *Product) GetCategoryID() string {
	return p.CategoryID.GetValue()
}

func (p *Product) GetQuantity() int {
	return p.Quantity.GetValue()
}

func (p *Product) GetPrice() float64 {
	return p.Price.GetValue()
}

func (p *Product) GetCreatedAt() time.Time {
	return p.CreatedAt
}

func (p *Product) GetUpdatedAt() time.Time {
	return p.UpdatedAt
}
