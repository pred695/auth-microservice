compose:
	docker-compose up
decompose:
	docker-compose down
build:
	sudo chmod -R 777 pgdata
	docker build -t auth-microservice .
connect:
	docker exec -it auth-microservice /bin/bash