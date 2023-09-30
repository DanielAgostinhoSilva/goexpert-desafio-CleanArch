package aws

import (
	"context"
	"github.com/DanielAgostinhoSilva/goexpert-desafio-CleanArch/src/infrastructure/env"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"log"
)

func NewSnsClient(env env.EnvConfig) *sns.Client {
	// Configuração da sessão AWS usando o AWS SDK for Go v2.
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithEndpointResolverWithOptions(
		aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
			return aws.Endpoint{
				URL:           env.AWSEndpoint,
				SigningRegion: region,
			}, nil
		}),
	))
	if err != nil {
		panic(err)
	}
	log.Println("SNS client initialized")
	return sns.NewFromConfig(cfg)
}
