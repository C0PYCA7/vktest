include config/.env
create:
	@echo "Создание миграции"
	goose -dir internal/storage/migrations create $(name) sql

up:
	@echo "Поднятие миграции"
	goose -dir internal/storage/migrations postgres $(DB_DSN) up

down:
	@echo "Откат миграции"
	goose -dir internal/storage/migrations postgres $(DB_DSN) down