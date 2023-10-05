package aws

import (
	"github.com/DanielAgostinhoSilva/goexpert-desafio-CleanArch/src/infrastructure/env"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

func NewSnsClient(env env.EnvConfig) *sns.SNS {
	cfg := aws.Config{
		Credentials: credentials.NewStaticCredentials("access_key_id", "secret_access_key", "session_token"),
		Endpoint:    aws.String("http://localhost:4566"), // URL do LocalStack SNS
		Region:      aws.String("us-east-1"),             // Região (não é usada com LocalStack)
	}

	sess := session.Must(session.NewSession(&cfg))
	return sns.New(sess)
}
