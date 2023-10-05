STACK_NAME=goexpert-desafio-cleanarch
#PROTO_GO_PATH=./../src/infrastructure/grpc
#PROTO_GO_OUT_PATH=./../src/infrastructure/grpc
#PROTO_PATH=./../src/infrastructure/grpc/protofiles
PROTO_GO_PATH=./../
PROTO_GO_OUT_PATH=./../
PROTO_PATH=./../src/infrastructure/grpc/protofiles

.PHONY: grpc-client

start:
	docker-compose -p ${STACK_NAME} up -d

stop:
	docker-compose -p ${STACK_NAME} stop

restart: stop start

clean:
	docker-compose -p ${STACK_NAME} down

logs:
	docker-compose -p ${STACK_NAME} logs -f

ps:
	docker-compose -p ${STACK_NAME} ps

grpc-generate:
	protoc --go_out=${PROTO_GO_PATH} --go-grpc_out=${PROTO_GO_OUT_PATH} --proto_path=${PROTO_PATH} order.proto

grpc-client:
	evans -r repl
