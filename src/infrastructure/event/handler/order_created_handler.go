package handler

import (
	"encoding/json"
	"fmt"
	"github.com/DanielAgostinhoSilva/goexpert-desafio-CleanArch/src/domain/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sns"
	"log"
	"os"
	"sync"
)

type OrderCreatedHandler struct {
	client *sns.SNS
}

func NewOrderCreatedHandler(client *sns.SNS) *OrderCreatedHandler {
	return &OrderCreatedHandler{client: client}
}

func (props *OrderCreatedHandler) Handle(event events.Event, wg *sync.WaitGroup) {
	defer wg.Done()
	log.Printf("Order created: %v\n", event.GetPayload())
	jsonOutput, _ := json.Marshal(event.GetPayload())

	result, err := props.client.Publish(&sns.PublishInput{
		TopicArn: aws.String("arn:aws:sns:us-east-1:000000000000:order-domain-event"), // Substitua pelo ARN correto
		Message:  aws.String(string(jsonOutput)),
	})

	if err != nil {
		fmt.Println("Erro ao publicar mensagem:", err)
		os.Exit(1)
	}

	log.Println("Message published successfully: ", result.MessageId)
}
