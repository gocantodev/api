.PHONY: server

run:
	cp .env.example .env
	docker-compose up
