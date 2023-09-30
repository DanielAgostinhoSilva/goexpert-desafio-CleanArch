package handler

import (
	"context"
	"encoding/json"
	"github.com/DanielAgostinhoSilva/goexpert-desafio-CleanArch/src/domain/events"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"log"
	"sync"
)

type OrderCreatedHandler struct {
	client *sns.Client
}

func (props *OrderCreatedHandler) Handle(event events.Event, wg *sync.WaitGroup) {
	defer wg.Done()
	log.Printf("Order created: %v\n", event.GetPayload())
	jsonOutput, _ := json.Marshal(event.GetPayload())

	input := &sns.PublishInput{
		Message:  aws.String(string(jsonOutput)),
		TopicArn: aws.String("arn:aws:sns:us-east-1:000000000000:order-domain-event"),
	}

	_, err := props.client.Publish(context.Background(), input)
	if err != nil {
		log.Println("failed to publish message, " + err.Error())
	}

	log.Println("Message published successfully!")
}
