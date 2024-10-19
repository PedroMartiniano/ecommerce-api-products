package vo

import "errors"

type Price struct {
	value float64
}

func NewPrice(value float64) (Price, error) {
	if value < 0.0 {
		return Price{}, errors.New("price cannot be negative")
	}

	return Price{
		value: value,
	}, nil
}

func (p *Price) GetValue() float64 {
	return p.value
}
