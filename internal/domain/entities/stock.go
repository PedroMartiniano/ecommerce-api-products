package entities

import (
	"time"

	"github.com/PedroMartiniano/ecommerce-api-products/internal/domain/vo"
)

type Stock struct {
	id        vo.UUID     `json:"id"`
	productID vo.UUID     `json:"product_id"`
	quantity  vo.Quantity `json:"quantity"`
	updatedAt time.Time   `json:"updated_at"`
}

// empty string for id if you're creating a new stock
func NewStock(id string, productID string, quantity int, updatedAt *time.Time) (Stock, error) {
	quantityVO, err := vo.NewQuantity(quantity)
	if err != nil {
		return Stock{}, err
	}

	if updatedAt == nil {
		updatedAt = new(time.Time)
	}

	return Stock{
		id:        vo.NewUUID(id),
		productID: vo.NewUUID(productID),
		quantity:  quantityVO,
		updatedAt: *updatedAt,
	}, nil
}

func (s *Stock) Add(quantity int) error {
	newValue := s.quantity.GetValue() + quantity
	newQuantity, err := vo.NewQuantity(newValue)
	if err != nil {
		return err
	}

	s.quantity = newQuantity
	s.updatedAt = time.Now()

	return nil
}

func (s *Stock) Subtract(quantity int) error {
	newValue := s.quantity.GetValue() - quantity
	newQuantity, err := vo.NewQuantity(newValue)
	if err != nil {
		return err
	}

	s.quantity = newQuantity
	s.updatedAt = time.Now()

	return nil
}

func (s *Stock) GetID() string {
	return s.id.GetValue()
}

func (s *Stock) GetProductID() string {
	return s.productID.GetValue()
}

func (s *Stock) GetQuantity() int {
	return s.quantity.GetValue()
}

func (s *Stock) GetUpdatedAt() time.Time {
	return s.updatedAt
}
