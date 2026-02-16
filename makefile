help:
	@echo "Available commands:"
	@echo "  run_setup_frontend  - Set up the frontend environment"
	@echo "  run_fe              - Start the frontend application"

run_setup_fe:
	@echo "Setting up frontend..."
	cd frontend && npm install
	@echo "Frontend setup complete."

run_fe:
	@echo "Starting frontend..."
	cd frontend && npm run dev