STACK_NAME=goexpert-desafio-cleanarch

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
