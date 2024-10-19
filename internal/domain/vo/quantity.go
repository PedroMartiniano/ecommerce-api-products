package vo

import "errors"

type Quantity struct {
	value int
}

func NewQuantity(value int) (Quantity, error) {
	if value < 0 {
		return Quantity{}, errors.New("unavailable quantity")
	}

	return Quantity{
		value: value,
	}, nil
}

func (q *Quantity) GetValue() int {
	return q.value
}
