package entities

import (
	"time"

	"github.com/PedroMartiniano/ecommerce-api-products/internal/domain/vo"
)

type Stock struct {
	ID        vo.UUID     `json:"id"`
	ProductID vo.UUID     `json:"product_id"`
	Quantity  vo.Quantity `json:"quantity"`
	UpdatedAt time.Time   `json:"updated_at"`
}

// empty string id if you're creating a new stock
func NewStock(id string, productID string, quantity int, updatedAt time.Time) (Stock, error) {
	quantityVO, err := vo.NewQuantity(quantity)
	if err != nil {
		return Stock{}, err
	}

	return Stock{
		ID:        vo.NewUUID(id),
		ProductID: vo.NewUUID(productID),
		Quantity:  quantityVO,
		UpdatedAt: updatedAt,
	}, nil
}

func (s *Stock) Add(quantity int) error {
	newValue := s.Quantity.GetValue() + quantity
	newQuantity, err := vo.NewQuantity(newValue)
	if err != nil {
		return err
	}

	s.Quantity = newQuantity
	s.UpdatedAt = time.Now()

	return nil
}

func (s *Stock) Subtract(quantity int) error {
	newValue := s.Quantity.GetValue() - quantity
	newQuantity, err := vo.NewQuantity(newValue)
	if err != nil {
		return err
	}

	s.Quantity = newQuantity
	s.UpdatedAt = time.Now()

	return nil
}
