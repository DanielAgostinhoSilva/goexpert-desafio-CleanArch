package main

import (
	application "github.com/DanielAgostinhoSilva/goexpert-desafio-CleanArch/src/application/order"
	"github.com/DanielAgostinhoSilva/goexpert-desafio-CleanArch/src/domain/events"
	event2 "github.com/DanielAgostinhoSilva/goexpert-desafio-CleanArch/src/domain/events/event"
	"github.com/DanielAgostinhoSilva/goexpert-desafio-CleanArch/src/domain/order"
	"github.com/DanielAgostinhoSilva/goexpert-desafio-CleanArch/src/infrastructure/aws"
	"github.com/DanielAgostinhoSilva/goexpert-desafio-CleanArch/src/infrastructure/database/mysql"
	"github.com/DanielAgostinhoSilva/goexpert-desafio-CleanArch/src/infrastructure/env"
	"github.com/DanielAgostinhoSilva/goexpert-desafio-CleanArch/src/infrastructure/event"
	"github.com/DanielAgostinhoSilva/goexpert-desafio-CleanArch/src/infrastructure/event/handler"
	"github.com/DanielAgostinhoSilva/goexpert-desafio-CleanArch/src/infrastructure/grpc"
	"github.com/DanielAgostinhoSilva/goexpert-desafio-CleanArch/src/infrastructure/webserver"
	"github.com/aws/aws-sdk-go/service/sns"
	"gorm.io/gorm"
	"sync"
)

var (
	config             *env.EnvConfig
	orderRepository    domain.OrderRepository
	db                 *gorm.DB
	snsClient          *sns.SNS
	eventDispatcher    events.EventDispatcher
	createOrderUseCase *application.CreateOrderUseCase
)

func init() {
	config = env.LoadConfig("./.env")
	mysql.MigrationUp(*config)
	db = mysql.Initialize(*config)
	orderRepository = mysql.NewOrderRepository(db)
	snsClient = aws.NewSnsClient(*config)
	eventDispatcher = event.NewEventDispatcher()
	eventDispatcher.Register("OrderCreated", handler.NewOrderCreatedHandler(snsClient))
	createOrderUseCase = application.NewCreateOrderUseCase(orderRepository, event2.NewOrderCreated(), eventDispatcher)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go webserver.Initialize(&wg, orderRepository, config.WebServerPort, createOrderUseCase)
	go grpc.Initialize(&wg, config.GRPCServerPort, createOrderUseCase, orderRepository)
	wg.Wait()
}
