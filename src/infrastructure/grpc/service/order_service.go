package service

import (
	"context"
	application "github.com/DanielAgostinhoSilva/goexpert-desafio-CleanArch/src/application/order"
	domain "github.com/DanielAgostinhoSilva/goexpert-desafio-CleanArch/src/domain/order"
	"github.com/DanielAgostinhoSilva/goexpert-desafio-CleanArch/src/infrastructure/grpc/pb"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	CreateOrderUseCase *application.CreateOrderUseCase
	OrderRepository    domain.OrderRepository
}

func NewOrderService(
	createOrderUseCase *application.CreateOrderUseCase,
	OrderRepository domain.OrderRepository,
) *OrderService {
	return &OrderService{
		CreateOrderUseCase: createOrderUseCase,
		OrderRepository:    OrderRepository,
	}
}

func (props *OrderService) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	dto := application.CreateOrderInput{
		ID:    in.Id,
		Price: float64(in.Price),
		Tax:   float64(in.Tax),
	}
	output, err := props.CreateOrderUseCase.Execute(dto)
	if err != nil {
		return nil, err
	}
	return &pb.CreateOrderResponse{
		Id:         output.ID,
		Price:      float32(output.Price),
		Tax:        float32(output.Tax),
		FinalPrice: float32(output.FinalPrice),
	}, nil
}

func (props *OrderService) ListOrder(context.Context, *pb.Blank) (*pb.OrderList, error) {
	orders, err := props.OrderRepository.FindAll()
	if err != nil {
		return nil, err
	}
	var ordersResponse []*pb.CreateOrderResponse
	for _, order := range orders {
		ordersResponse = append(ordersResponse, &pb.CreateOrderResponse{
			Id:         order.ID,
			Price:      float32(order.Price),
			Tax:        float32(order.Tax),
			FinalPrice: float32(order.FinalPrice),
		})
	}

	return &pb.OrderList{
		Orders: ordersResponse,
	}, nil
}
