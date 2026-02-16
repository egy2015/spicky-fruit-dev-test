# spicky-fruit-dev-test

This is a fullstack internal dashboard application for monitoring incoming payments.  
The backend is built using Golang, and the frontend is built using Vue 3.

The application provides authentication, protected dashboard routes, and payment data visualization based on REST APIs defined in an OpenAPI specification.

---

## Environment Requirements

This project is tested and expected to run in the following environment:

```bash
Go 1.21 or newer
Node.js 20 or newer
Docker installed
Docker Compose installed
Make installed
macOS (tested on MacBook Pro)
```

---

## Project Structure

```text
.
├── backend/
├── frontend/
├── openapi.yaml
├── Makefile
└── README.md
```

---

## Backend

The backend service is implemented using Golang and exposes RESTful APIs defined in `openapi.yaml`.

### Run Backend Server (Local)

```bash
cd backend
make run
```

If Makefile is not available:

```bash
go mod tidy
go run main.go
```

### Run Backend Server (Production Build)

```bash
cd backend
make build
make start
```

---

## Frontend

The frontend application is built using Vue 3, Vite, Pinia, and Vue Router.

### Install Frontend Dependencies

```bash
make run_setup_fe
```

### Run Frontend (Local)

```bash
make run_fe
```

The frontend will be available on the default Vite development port.

### Build Frontend (Production)

```bash
make build_fe
```

---

## API Documentation

The OpenAPI documentation is available in the root of the repository:

```bash
openapi.yaml
```

You can inspect available endpoints and request/response schemas from this file.

---

## Authentication Flow

- Authentication uses JWT.
- Users must log in before accessing the dashboard.
- Protected routes require a valid token.

---

## Accessing the Application

After running both backend and frontend services, access the application at:

```bash
http://localhost:5173/login
```

---

## Notes

- Dashboard routes are protected.
- Payments data is fetched from backend REST APIs.
- API base URL is configured via frontend environment variables.

---

## Evidence

Optional: attach video or screen recording showing the application running.

---

For more detailed information:
- Backend documentation: [backend/README.md](backend/README.md)
- Frontend documentation: [frontend/README.md](frontend/README.md)
