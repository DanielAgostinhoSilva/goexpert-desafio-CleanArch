package webserver

import (
	"encoding/json"
	application "github.com/DanielAgostinhoSilva/goexpert-desafio-CleanArch/src/application/order"
	"github.com/DanielAgostinhoSilva/goexpert-desafio-CleanArch/src/domain/order"
	"net/http"
)

type OrderHandler struct {
	OrderRepository domain.OrderRepository
}

func NewOrderHandler(OrderRepository domain.OrderRepository) *OrderHandler {
	return &OrderHandler{
		OrderRepository: OrderRepository,
	}
}

func (props *OrderHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input application.CreateOrderInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	createOrder := application.NewCreateOrderUseCase(props.OrderRepository)
	output, err := createOrder.Execute(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
