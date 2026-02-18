# spicky-fruit-dev-test

This is a fullstack internal dashboard application for monitoring incoming payments.  
The backend is built using Golang, and the frontend is built using Vue 3.

The application provides authentication, protected dashboard routes, and payment data visualization based on REST APIs defined in an OpenAPI specification.

---

## Environment Requirements

This project is tested and expected to run in the following environment:

```bash
Go 1.25.5
Node.js 20.19.0
Docker installed
Docker Compose installed
Make installed
macOS (tested on MacBook Pro)
```


To avoid version conflicts, it is strongly recommended to use version managers:
- **Go**: [GVM (Go Version Manager)](https://github.com/moovweb/gvm)
- **Node.js**: [NVM (Node Version Manager)](https://github.com/nvm-sh/nvm)

---

## Project Structure

```text
.
├── backend/
├── frontend/
├── docker-compose.yaml
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

---

## Makefile Commands

### Frontend
- `make run_setup_fe` : Set up frontend dependencies
- `make run_fe`       : Start frontend development server ([http://localhost:5173](http://localhost:5173))

### Backend
- `make run_setup_be` : Set up environment, vendor, OpenAPI & JWT secret
- `make run_be`       : Start backend server ([http://localhost:8080](http://localhost:8080))

### Docker
- `make docker-build`   : Build images (clean build)
- `make docker-restart` : Restart all containers
- `make docker-status`  : Check Docker container health/status

---

## API Overview

**Base Path**: `/dashboard/v1`

| Method | Endpoint | Description |
| :--- | :--- | :--- |
| `POST` | `/auth/login` | Login and get JWT token |
| `GET` | `/payments` | Get list of payments (Auth Required) |
| `PUT` | `/payments/:id` | Update payment (Merchant & Amount) |

---

## Frontend Authentication

**Login Page**: [http://localhost:5173/login](http://localhost:5173/login)

**Default Test Users:**
- `cs@test.com` / `password`
- `operation@test.com` / `password`

---

## Notes
- **CORS**: Enabled for `http://localhost:5173` for local development.
- **Scope**: Features follow take-home test requirements strictly to avoid overengineering.

## Evidence
Please include the following in your submission:
1. Screenshots or screen recordings of the application.
2. Status of Docker running services (`make docker-status`).
3. Evidence of successful login and payment update flow.

---
*See also:* [backend/README.md](./backend/README.md) | [frontend/README.md](./frontend/README.md)

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
