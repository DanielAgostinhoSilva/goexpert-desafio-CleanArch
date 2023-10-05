package application

import (
	"github.com/DanielAgostinhoSilva/goexpert-desafio-CleanArch/src/domain/events"
	domain "github.com/DanielAgostinhoSilva/goexpert-desafio-CleanArch/src/domain/order"
)

type CreateOrderInput struct {
	ID    string
	Price float64
	Tax   float64
}

type CreateOrderOutput struct {
	ID         string
	Price      float64
	Tax        float64
	FinalPrice float64
}

type CreateOrderUseCase struct {
	OrderRepository domain.OrderRepository
	OrderCreated    events.Event
	EventDispatcher events.EventDispatcher
}

func NewCreateOrderUseCase(
	OrderRepository domain.OrderRepository,
	OrderCreated events.Event,
	EventDispatcher events.EventDispatcher,
) *CreateOrderUseCase {
	return &CreateOrderUseCase{
		OrderRepository: OrderRepository,
		OrderCreated:    OrderCreated,
		EventDispatcher: EventDispatcher,
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

	props.OrderCreated.SetPayload(createOrderOutput)
	props.EventDispatcher.Dispatch(props.OrderCreated)

	return createOrderOutput, nil
}
