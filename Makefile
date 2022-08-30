.PHONY: gocanto\:server

include .env

CURRENT_DIR = $(shell pwd)
DB_MIGRATIONS_DIR = database/migrations

clean:
	docker container prune -f
	docker image prune -f
	docker volume prune -f

migrate\:up:
	docker run -v $(CURRENT_DIR)/$(DB_MIGRATIONS_DIR):/$(DB_MIGRATIONS_DIR) --network server migrate/migrate -path=/$(DB_MIGRATIONS_DIR)/ -database ${POSTGRES_URL} up 1

migrate\:down:
	docker run -v $(CURRENT_DIR)/$(DB_MIGRATIONS_DIR):/$(DB_MIGRATIONS_DIR) --network server migrate/migrate -path=/$(DB_MIGRATIONS_DIR)/ -database ${POSTGRES_URL} down 1
