# Variables
APP_NAME=collector-service

# Tasks
default: run

run:
	@go run cmd/${APP_NAME}/main.go
run-docker-compose:
	@docker compose -f build/docker-compose.yml up --build -d
run-build-docker:
	@docker run -p 3000:3000 ${APP_NAME}
build:
	@go build -o ${APP_NAME} cmd/${APP_NAME}/main.go
build-docker:
	@docker build -f build/Dockerfile . -t ${APP_NAME}