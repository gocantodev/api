.PHONY: gocanto\:server

include .env

CURRENT_DIR = $(shell pwd)
DB_MIGRATIONS_DIR = database/migrations

test:
	echo ${POSTGRES_VERSION}

migrate\:up:
	docker run -v $(CURRENT_DIR)/$(DB_MIGRATIONS_DIR):/$(DB_MIGRATIONS_DIR) --network host migrate/migrate -path=/$(DB_MIGRATIONS_DIR)/ -database ${POSTGRES_URL} up 1

migrate\:down:
	docker run -v $(CURRENT_DIR)/$(DB_MIGRATIONS_DIR):/$(DB_MIGRATIONS_DIR) --network host migrate/migrate -path=/$(DB_MIGRATIONS_DIR)/ -database ${POSTGRES_URL} down 1
