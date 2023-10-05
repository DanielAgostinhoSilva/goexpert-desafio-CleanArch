package grpc

import (
	"fmt"
	application "github.com/DanielAgostinhoSilva/goexpert-desafio-CleanArch/src/application/order"
	domain "github.com/DanielAgostinhoSilva/goexpert-desafio-CleanArch/src/domain/order"
	"github.com/DanielAgostinhoSilva/goexpert-desafio-CleanArch/src/infrastructure/grpc/pb"
	"github.com/DanielAgostinhoSilva/goexpert-desafio-CleanArch/src/infrastructure/grpc/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"sync"
)

func Initialize(
	wg *sync.WaitGroup,
	port string,
	createOrderUseCase *application.CreateOrderUseCase,
	orderRepository domain.OrderRepository,
) {
	defer wg.Done()
	grpcServer := grpc.NewServer()
	orderService := service.NewOrderService(createOrderUseCase, orderRepository)
	pb.RegisterOrderServiceServer(grpcServer, orderService)
	reflection.Register(grpcServer)

	log.Println("Starting gRPC server on port", port)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		panic(err)
	}
	grpcServer.Serve(lis)
}
