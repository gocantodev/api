.PHONY: server

serve-app:
	cp .env.example .env
	docker-compose up -d

down:
	docker-compose down

migrate-up:
	docker run -v /Users/gus/Sites/server/database/migrations:/database/migrations --network host migrate/migrate -path=/database/migrations/ -database postgres://root:root@localhost:5432/gocanto?sslmode=disable up 1

migrate-down:
	docker run -v /Users/gus/Sites/server/database/migrations:/database/migrations --network host migrate/migrate -path=/database/migrations/ -database postgres://root:root@localhost:5432/gocanto?sslmode=disable down 1
