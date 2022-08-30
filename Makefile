.PHONY: gocanto\:server

include .env

APP_NETWORK = gocanto
APP_PATH = $(shell pwd)
APP_MIGRATIONS = database/migrations

serve:
	docker compose up --wait

stop:
	docker compose down

prune:
	docker compose down --remove-orphans
	docker container prune -f
	docker image prune -f
	docker volume prune -f
	docker network prune -f

prune\:data:
	@rm -rf $(APP_PATH)/database/data/*

migrate\:up:
	$(call migrate) up 1

migrate\:down:
	$(call migrate) down 1

define migrate
	@# param 1 is the method to be called.
	@# param 2 is the number of batches to excecute.

	docker run -v $(APP_PATH)/$(APP_MIGRATIONS):/$(APP_MIGRATIONS) --network ${APP_NETWORK} migrate/migrate -path=/$(APP_MIGRATIONS)/ -database ${POSTGRES_URL} $(1) $(2)
endef
