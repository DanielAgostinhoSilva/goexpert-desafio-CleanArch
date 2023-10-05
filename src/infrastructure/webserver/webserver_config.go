package webserver

import (
	application "github.com/DanielAgostinhoSilva/goexpert-desafio-CleanArch/src/application/order"
	domain "github.com/DanielAgostinhoSilva/goexpert-desafio-CleanArch/src/domain/order"
	"github.com/DanielAgostinhoSilva/goexpert-desafio-CleanArch/src/infrastructure/webserver/controller"
	"log"
	"sync"
)

func Initialize(
	wg *sync.WaitGroup,
	repository domain.OrderRepository,
	port string,
	createOrderUseCase *application.CreateOrderUseCase,
) {
	defer wg.Done()
	webserver := NewWebServer(port)
	webserver.AddController(controller.NewOrderController(repository, createOrderUseCase))
	log.Println("Starting web server on port ", port)
	webserver.Start()
}
