package webserver

import (
	domain "github.com/DanielAgostinhoSilva/goexpert-desafio-CleanArch/src/domain/order"
	"log"
	"sync"
)

func Initialize(wg *sync.WaitGroup, repository domain.OrderRepository, port string) {
	defer wg.Done()
	webserver := NewWebServer(port)
	orderHandler := NewOrderHandler(repository)
	webserver.AddHandler("POST", "/orders", orderHandler.Create)
	log.Println("Starting web server on port ", port)
	webserver.Start()
}