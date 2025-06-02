# Run tests with verbose output
test:
	@go test -v ./...

# Start the app using Docker Compose
start:
	docker compose up --build

# Tear down Docker containers
stop:
	docker compose down