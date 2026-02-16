# spicky-fruit-dev-test



## Environment Requirements

This project is tested and expected to run in the following environment:

- Go 1.21
- Node.js 20
- Docker installed
- Docker Compose installed
- Make installed
- and a Macbook Pro M1


## Project Structure

```text
.
├── backend/
├── frontend/
├── openapi.yaml
├── Makefile
└── README.md
```

## Frontend

The frontend is built using Vue 3, Vite, and Pinia.

### Install Frontend Dependencies

Frontend dependencies are installed automatically using Make:

```bash
make run_setup_fe
```

This command installs all required frontend dependencies.

### Run Frontend Development Server

```bash
make run_fe
```

The frontend will be available on the default Vite development port.

### Build Frontend for Production

```bash
make build_fe
```

## Backend

Backend setup and execution instructions are available in the `backend` directory.

## Notes

- Authentication uses JWT.
- Dashboard routes are protected and require a valid token.
- Payments data is fetched from backend REST APIs defined in `openapi.yaml`.
