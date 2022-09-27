package entity

import "errors"

type Order struct {
	ID         string
	Price      float64
	Tax        float64
	FinalPrice float64
}

func NewOrder(id string, price float64, tax float64) (*Order, error) {
	newOrder := &Order{
		ID:    id,
		Price: price,
		Tax:   tax,
	}
	err := newOrder.IsValid()
	if err != nil {
		return nil, err
	}
	newOrder.FinalPrice = price + (price * (tax / 100))
	return newOrder, nil
}

func (o *Order) IsValid() error {
	if o.ID == "" {
		return errors.New("invalid id")
	}
	if o.Price == 0.0 {
		return errors.New("invalid price")
	}
	if o.Tax == 0 {
		return errors.New("invalid tax")
	}
	return nil
}
