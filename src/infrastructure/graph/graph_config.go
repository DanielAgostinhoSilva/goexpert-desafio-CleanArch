package graph

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	application "github.com/DanielAgostinhoSilva/goexpert-desafio-CleanArch/src/application/order"
	domain "github.com/DanielAgostinhoSilva/goexpert-desafio-CleanArch/src/domain/order"
	"log"
	"net/http"
	"sync"
)

func Initialize(
	wg *sync.WaitGroup,
	port string,
	createOrderUseCase *application.CreateOrderUseCase,
	orderRepository domain.OrderRepository,
) {
	defer wg.Done()
	srv := handler.NewDefaultServer(NewExecutableSchema(Config{Resolvers: &Resolver{
		OrderRepository:    orderRepository,
		CreateOrderUseCase: createOrderUseCase,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Println("Starting GraphQL server on port", port)
	log.Printf("Connect to http://localhost:%s/ for GraphQL playground\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
