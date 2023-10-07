# goexpert-desafio-CleanArch

## Desafio Clean Architecture
Olá devs!

Agora é a hora de botar a mão na massa. Pra este desafio, você precisará criar a listagem das orders.

Esta listagem precisa ser feita com:

- Endpoint REST (GET /order)

- Service ListOrders com GRPC

- Query ListOrders GraphQL

Não esqueça de criar as migrações necessárias e o arquivo api.http com a request para criar e listar as orders.

## Visão geral
- Endpoint REST
  - POST /orders "cria um order"
  - GET /orders "busca todos os orders na base de dados"

- Service GRPC
  - call CreateOrder "cria um order"
  - call ListOrder "busca todos os orders na base de dados"

- Service GraphQL
  - mutation createOrder "cria um order"
  - query orders "busca todos os orders na base de dados"

Agumas tecnonologias foi trocada para propor um desafio paracido com que uso no meu dia a dia.

RabbitMQ foi trcado por SNS


Este é um layout do projeto foi ajustado conforme a visao do projeto da galera do [Colibri Project](https://colibriproject.io/) que gostei muito da organização

### `/dev`

Onde fica todos os arquivo de desenvolvimento local

### `/src`

Diretorio onde fica todos os fontes do projeto.

### `/test`

Diretorio onde fica todos os testes unitarios e de integração.
