package entity

import "errors"

type Order struct {
	ID         string
	Price      float64
	Tax        float64
	FinalPrice float64
}

type OrderRepositoryInterface interface {
	Save(order *Order) error
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
	return newOrder, nil
}

func (o *Order) CalculateFinalPrice() error {
	o.FinalPrice = o.Price + o.Tax
	return o.IsValid()
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
