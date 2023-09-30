package webserver

import (
	domain "github.com/DanielAgostinhoSilva/goexpert-desafio-CleanArch/src/domain/order"
	"github.com/DanielAgostinhoSilva/goexpert-desafio-CleanArch/src/infrastructure/webserver/controller"
	"log"
	"sync"
)

func Initialize(wg *sync.WaitGroup, repository domain.OrderRepository, port string) {
	defer wg.Done()
	webserver := NewWebServer(port)
	webserver.AddController(controller.NewOrderController(repository))
	log.Println("Starting web server on port ", port)
	webserver.Start()
}
