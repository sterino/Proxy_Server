build:
		docker build .

up:
		docker-compose up

down:
		docker-compose down

restart: down up

.PHONY: build up down restart