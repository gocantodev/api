.PHONY: server

up:
	cp .env.example .env
	docker-compose up

down:
	docker-compose down
