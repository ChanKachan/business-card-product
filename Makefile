generate-proto-card-product-v1:
	@mkdir -p ./pb/cardProduct
	@protoc -I ./proto/cardProduct \
        --go_out=./pb/cardProduct  --go_opt=paths=source_relative \
        --go-grpc_out=./pb/cardProduct  --go-grpc_opt=paths=source_relative \
        proto/cardProduct/*.proto
	@echo 'Success Generate Proto Card Product'

# Переменные для Docker
COMPOSE = docker compose --env-file ./.env -f ./docker-compose.yml
APP_SERVICE = app

# Запуск контейнеров
start:
	@$(COMPOSE) up -d

# Остановка контейнеров
stop:
	@$(COMPOSE) down

# Рестарт контейнеров
restart:
	@$(COMPOSE) up -d --build --no-deps --force-recreate $(APP_SERVICE)
