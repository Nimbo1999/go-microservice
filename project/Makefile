BROKER_BINARY=brokerApp
AUTH_BINARY=authenticationApp

up:
	@echo "Starting Docker images..."
	docker-compose up -d
	@echo "Docker images started!"

up_build:
	@echo "Stopping docker images (if running...)"
	docker-compose down
	@echo "Building (when required) and starting docker images..."
	docker-compose up --build -d
	@echo "Docker images built and started!"

down:
	@echo "Stopping docker compose..."
	docker-compose down
	@echo "Done!"

build_broker:
	@echo "Building broker binary..."
	cd ../broker-service ** env GOOS=linux CGO_ENABLED=0 go build -o ${BROKER_BINARY} ./cmd/api
	@echo "Done!"

build_front:
	@echo "Building frontend"
	cd ../frontend && yarn build
	@echo "Done!"

build_auth:
	@echo "Building auth binary..."
	cd ../authentication-service ** env GOOS=linux CGO_ENABLED=0 go build -o ${AUTH_BINARY} ./cmd/api
	@echo "Done!"

start: build_front
	@echo "Starting frontend"
	cd ../frontend/dist && serve .
	@echo ""

# Migrate databases wit
# goose postgres "postgresql://postgres:password@localhost/users?sslmode=disable&connect_timeout=5" up
# GOOSE_DRIVER=postgres GOOSE_DBSTRING="postgresql://postgres:password@localhost/users?sslmode=disable&connect_timeout=5" goose up
