package graph

import (
	application "github.com/DanielAgostinhoSilva/goexpert-desafio-CleanArch/src/application/order"
	domain "github.com/DanielAgostinhoSilva/goexpert-desafio-CleanArch/src/domain/order"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	CreateOrderUseCase *application.CreateOrderUseCase
	OrderRepository    domain.OrderRepository
}
