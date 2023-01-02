.PHONY: gocanto\:api

include .env

DB_NETWORK = gocanto
APP_PATH = $(shell pwd)
DB_MIGRATIONS_PATH = database/migrations

run\:api:
	cd cmd/api && go mod tidy && go mod download && \
	CGO_ENABLED=0 go run -tags api github.com/gocantodev/api/cmd/api

up:
	docker compose up nginx

build:
	docker compose up nginx --build

build\:fresh:
	make docker\:prune && \
	docker compose up nginx --build

down:
	docker compose down

docker\:prune:
	docker compose down --remove-orphans
	docker container prune -f
	docker image prune -f
	docker volume prune -f
	docker network prune -f

db\:reset:
	make docker\:prune
	make up
	@rm -rf $(APP_PATH)/database/data/*

db\:migrate\:up:
	$(call migrate) up 1

db\:migrate\:down:
	$(call migrate) down 1

define migrate
	@# param 1 is the method to be called.
	@# param 2 is the number of batches to excecute.

	docker run -v $(APP_PATH)/$(DB_MIGRATIONS_PATH):/$(DB_MIGRATIONS_PATH) --network ${DB_NETWORK} migrate/migrate -path=/$(DB_MIGRATIONS_PATH)/ -database ${POSTGRES_URL} $(1) $(2)
endef
