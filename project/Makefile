BROKER_BINARY=brokerApp
AUTH_BINARY=authenticationApp
LOGGER_BINARY=loggerApp
MAIL_BINARY=mailApp

up:
	@echo "Starting Docker images..."
	docker-compose up -d
	@echo "Docker images started!"
	@echo "Initializing the migration..."
	make migrate
	@echo "Migrations completed!"

up_build:
	@echo "Stopping docker images (if running...)"
	docker-compose down
	@echo "Building (when required) and starting docker images..."
	docker-compose up --build -d
	@echo "Docker images built and started!"
	@echo "Initializing the migration..."
	make migrate
	@echo "Migrations completed!"

down:
	@echo "Stopping docker compose..."
	docker-compose down
	@echo "Done!"

build_broker:
	@echo "Building broker binary..."
	cd ../broker-service && GOOS=linux CGO_ENABLED=0 go build -o ${BROKER_BINARY} ./cmd/api
	@echo "Done!"

build_front:
	@echo "Building frontend"
	cd ../frontend && yarn build
	@echo "Done!"

build_auth:
	@echo "Building auth binary..."
	cd ../authentication-service && GOOS=linux CGO_ENABLED=0 go build -o ${AUTH_BINARY} ./cmd/api
	@echo "Done!"

build_logger:
	@echo "Building auth binary..."
	cd ../logger-service && GOOS=linux CGO_ENABLED=0 go build -o ${LOGGER_BINARY} ./cmd/api
	@echo "Done!"

build_mail:
	@echo "Building auth binary..."
	cd ../mail-service && GOOS=linux CGO_ENABLED=0 go build -o ${MAIL_BINARY} ./cmd/api
	@echo "Done!"

build_all_macos:
	@echo "Building auth binary..."
	cd ../broker-service && GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -o ${BROKER_BINARY} ./cmd/api
	cd ../authentication-service && GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -o ${AUTH_BINARY} ./cmd/api
	cd ../logger-service && GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -o ${LOGGER_BINARY} ./cmd/api
	cd ../mail-service && GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -o ${MAIL_BINARY} ./cmd/api
	@echo "Done!"

build_all_linux:
	@echo "Building auth binary..."
	cd ../broker-service && GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ${BROKER_BINARY} ./cmd/api
	cd ../authentication-service && GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ${AUTH_BINARY} ./cmd/api
	cd ../logger-service && GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ${LOGGER_BINARY} ./cmd/api
	cd ../mail-service && GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ${MAIL_BINARY} ./cmd/api
	@echo "Done!"

start: build_front
	@echo "Starting frontend"
	cd ../frontend/dist && serve .
	@echo ""

migrate:
	cd ../authentication-service/migrations && goose postgres "postgresql://postgres:password@localhost/users?sslmode=disable&connect_timeout=5" up

# Migrate databases
# goose postgres "postgresql://postgres:password@localhost/users?sslmode=disable&connect_timeout=5" up
# GOOSE_DRIVER=postgres GOOSE_DBSTRING="postgresql://postgres:password@localhost/users?sslmode=disable&connect_timeout=5" goose up
