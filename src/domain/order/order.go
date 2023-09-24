package domain

import "errors"

type Order struct {
	ID         string
	Price      float64
	Tax        float64
	FinalPrice float64
}

func NewOrder(id string, price float64, tax float64) (*Order, error) {
	order := &Order{
		ID:    id,
		Price: price,
		Tax:   tax,
	}
	err := order.IsValid()
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (props *Order) IsValid() error {
	if props.ID == "" {
		return errors.New("invalid id")
	}
	if props.Price <= 0 {
		return errors.New("invalid price")
	}
	if props.Tax <= 0 {
		return errors.New("invalid tax")
	}
	return nil
}

func (props *Order) CalculateFinalPrice() error {
	props.FinalPrice = props.Price + props.Tax
	err := props.IsValid()
	if err != nil {
		return err
	}
	return nil
}
