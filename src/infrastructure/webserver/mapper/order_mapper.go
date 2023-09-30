package mapper

import (
	application "github.com/DanielAgostinhoSilva/goexpert-desafio-CleanArch/src/application/order"
	domain "github.com/DanielAgostinhoSilva/goexpert-desafio-CleanArch/src/domain/order"
	"github.com/DanielAgostinhoSilva/goexpert-desafio-CleanArch/src/infrastructure/webserver/model"
)

type OrderMapper struct {
}

func NewOrderMapper() *OrderMapper {
	return &OrderMapper{}
}

func (props *OrderMapper) OrderInputToCreateOrderInput(orderInput model.OrderInput) application.CreateOrderInput {
	return application.CreateOrderInput{
		ID:    orderInput.ID,
		Price: orderInput.Price,
		Tax:   orderInput.Tax,
	}
}

func (props *OrderMapper) CreateOrderOutputToOrderModel(orderInput application.CreateOrderOutput) *model.OrderModel {
	return &model.OrderModel{
		ID:    orderInput.ID,
		Price: orderInput.Price,
		Tax:   orderInput.Tax,
	}
}

func (props *OrderMapper) OrderToOrderModel(order domain.Order) *model.OrderModel {
	return &model.OrderModel{
		ID:         order.ID,
		Price:      order.Price,
		Tax:        order.Tax,
		FinalPrice: order.FinalPrice,
	}
}

func (props *OrderMapper) OrdersToCollectionOrderModel(orders []domain.Order) []*model.OrderModel {
	var ordersModel []*model.OrderModel
	for _, order := range orders {
		ordersModel = append(ordersModel, props.OrderToOrderModel(order))
	}
	return ordersModel
}
