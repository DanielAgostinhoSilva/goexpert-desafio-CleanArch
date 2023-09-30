package main

import (
	domain "github.com/DanielAgostinhoSilva/goexpert-desafio-CleanArch/src/domain/order"
	"github.com/DanielAgostinhoSilva/goexpert-desafio-CleanArch/src/infrastructure/aws"
	"github.com/DanielAgostinhoSilva/goexpert-desafio-CleanArch/src/infrastructure/database/mysql"
	"github.com/DanielAgostinhoSilva/goexpert-desafio-CleanArch/src/infrastructure/env"
	"github.com/DanielAgostinhoSilva/goexpert-desafio-CleanArch/src/infrastructure/webserver"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"gorm.io/gorm"
	"sync"
)

var (
	config          *env.EnvConfig
	orderRepository domain.OrderRepository
	db              *gorm.DB
	snsClient       *sns.Client
)

func init() {
	config = env.LoadConfig("./.env")
	mysql.MigrationUp(*config)
	db = mysql.Initialize(*config)
	orderRepository = mysql.NewOrderRepository(db)
	snsClient = aws.NewSnsClient(*config)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go webserver.Initialize(&wg, orderRepository, config.WebServerPort)
	wg.Wait()
}
