help:
	@echo "Available commands:"
	@echo "  run_setup_frontend  - Set up the frontend environment"
	@echo "  run_fe              - Start the frontend application"
	@echo "  docker-build        - Build the Docker images"
	@echo "  docker-restart      - Restart the Docker containers"
	@echo "  docker-status       - Check the status of Docker containers"

#backend
run_setup_be:
	@echo "Setting up backend..."
	@echo "Tidying and vendoring Go modules... (this may take a moment)"
	@echo "Make sure you have Go installed (go1.25.5 Recommended) and properly set up in your environment."
	cd backend && go mod tidy
	cd backend && go mod vendor
	cd backend && cp env.sample .env
	cd backend && make tool-openapi
	cd backend && make openapi-gen
	cd backend && make dep
	cd backend && make gen-secret
	@echo "Backend setup complete."

run_be:
	@echo "Starting backend..."
	cd backend && make run

#frontend
run_setup_fe:
	@echo "Setting up frontend..."
	@echo "Make sure you have Node.js installed (v18 or later like v20.19.0  recommended) and npm properly set up in your environment."
	cd frontend && npm install
	@echo "Frontend setup complete."

run_fe:
	@echo "Starting frontend..."
	cd frontend && npm run dev


#docker
docker-build:
	docker compose down -v
	docker compose build --no-cache

docker-restart:
	docker compose down -v
	docker compose up -d

docker-status:
	docker ps