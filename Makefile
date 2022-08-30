.PHONY: gocanto\:server

include .env

API_NETWORK = gocanto
CURRENT_DIR = $(shell pwd)
DB_MIGRATIONS_DIR = database/migrations

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
	@rm -rf $(CURRENT_DIR)/database/data/*

migrate\:up:
	docker run -v $(CURRENT_DIR)/$(DB_MIGRATIONS_DIR):/$(DB_MIGRATIONS_DIR) --network ${API_NETWORK} migrate/migrate -path=/$(DB_MIGRATIONS_DIR)/ -database ${POSTGRES_URL} up 1

migrate\:down:
	docker run -v $(CURRENT_DIR)/$(DB_MIGRATIONS_DIR):/$(DB_MIGRATIONS_DIR) --network ${API_NETWORK} migrate/migrate -path=/$(DB_MIGRATIONS_DIR)/ -database ${POSTGRES_URL} down 1
