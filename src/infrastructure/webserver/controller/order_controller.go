package controller

import (
	"encoding/json"
	application "github.com/DanielAgostinhoSilva/goexpert-desafio-CleanArch/src/application/order"
	domain "github.com/DanielAgostinhoSilva/goexpert-desafio-CleanArch/src/domain/order"
	"github.com/DanielAgostinhoSilva/goexpert-desafio-CleanArch/src/infrastructure/webserver/mapper"
	"github.com/DanielAgostinhoSilva/goexpert-desafio-CleanArch/src/infrastructure/webserver/model"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type OrderHandler struct {
	OrderRepository domain.OrderRepository
}

func NewOrderController(OrderRepository domain.OrderRepository) *OrderHandler {
	return &OrderHandler{
		OrderRepository: OrderRepository,
	}
}

func (props *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	mapper := mapper.NewOrderMapper()
	var orderInput model.OrderInput

	err := json.NewDecoder(r.Body).Decode(&orderInput)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createOrder := application.NewCreateOrderUseCase(props.OrderRepository)
	createOrderOutput, err := createOrder.Execute(mapper.OrderInputToCreateOrderInput(orderInput))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(mapper.CreateOrderOutputToOrderModel(createOrderOutput))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (props *OrderHandler) GetAllOrders(w http.ResponseWriter, r *http.Request) {
	mapper := mapper.NewOrderMapper()
	orders, err := props.OrderRepository.FindAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ordersModel := mapper.OrdersToCollectionOrderModel(orders)
	err = json.NewEncoder(w).Encode(ordersModel)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (props *OrderHandler) Router(router chi.Router) {
	router.Post("/", props.CreateOrder)
	router.Get("/", props.GetAllOrders)
}
func (props *OrderHandler) Path() string {
	return "/orders"
}
