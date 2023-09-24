package application

import domain "github.com/DanielAgostinhoSilva/goexpert-desafio-CleanArch/src/domain/order"

type CreateOrderInput struct {
	ID    string  `json:"id"`
	Price float64 `json:"price"`
	Tax   float64 `json:"tax"`
}

type CreateOrderOutput struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
}

type CreateOrderUseCase struct {
	OrderRepository domain.OrderRepository
}

func NewCreateOrderUseCase(orderRepository domain.OrderRepository) *CreateOrderUseCase {
	return &CreateOrderUseCase{
		OrderRepository: orderRepository,
	}
}

func (props *CreateOrderUseCase) Execute(input CreateOrderInput) (CreateOrderOutput, error) {
	order := domain.Order{
		ID:    input.ID,
		Price: input.Price,
		Tax:   input.Tax,
	}
	order.CalculateFinalPrice()
	if err := props.OrderRepository.Save(&order); err != nil {
		return CreateOrderOutput{}, err
	}

	createOrderOutput := CreateOrderOutput{
		ID:         order.ID,
		Price:      order.Price,
		Tax:        order.Tax,
		FinalPrice: order.Price + order.Tax,
	}

	return createOrderOutput, nil
}
