.PHONY: gocanto\:server

include .env

DB_NETWORK = gocanto
APP_PATH = $(shell pwd)
DB_MIGRATIONS_PATH = database/migrations

api\:build:
	docker compose build server

api\:run:
	docker compose run server

api\:fresh:
	make api\:build
	make api\:run

db\:up:
	docker compose up --wait

db\:down:
	docker compose down

db\:prune:
	docker compose down --remove-orphans
	docker container prune -f
	docker image prune -f
	docker volume prune -f
	docker network prune -f

db\:reset: db\:prune
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
