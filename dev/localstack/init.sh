# Criando um topico order-domain-event
awslocal sns create-topic --name order-domain-event

# Criando uma fila order-created
awslocal sqs create-queue --queue-name order-created

# Criando uma assinatura ao topico order-domain-event
awslocal sns subscribe --topic-arn arn:aws:sns:us-east-1:000000000000:order-domain-event --protocol sqs --notification-endpoint arn:aws:sqs:us-east-1:000000000000:order-created --attributes RawMessageDelivery=true

echo "localstack initialize finished"
