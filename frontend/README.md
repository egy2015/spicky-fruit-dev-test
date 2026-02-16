# Payments Dashboard Frontend

A Vue 3 + Vite frontend for an internal payments monitoring dashboard.

## Tech Stack

- **Vue 3** - Progressive JavaScript framework
- **Vite 5** - Next-generation build tool
- **Pinia** - State management for Vue
- **Vue Router 4** - Client-side routing
- **Axios** - HTTP client (optional, using fetch in API service)

## Project Structure

```
src/
├── main.js                 # Application entry point
├── App.vue                 # Root component
├── router/
│   └── index.js           # Router configuration and guards
├── stores/
│   └── auth.js            # Pinia auth store
├── services/
│   └── api.js             # API service layer
├── views/
│   ├── auth/
│   │   └── LoginView.vue  # Login page
│   ├── dashboard/
│   │   └── DashboardView.vue # Protected dashboard
│   └── NotFoundView.vue   # 404 page
├── components/
│   ├── layout/
│   │   └── DashboardLayout.vue
│   ├── payments/
│   │   ├── PaymentsTable.vue
│   │   ├── PaymentRow.vue
│   │   └── StatusBadge.vue
│   └── common/
│       └── Loading.vue
└── assets/
    └── styles/
        └── main.css
```

## Features

- ✅ JWT-based authentication
- ✅ Protected routes with route guards
- ✅ Login/logout functionality
- ✅ Payments dashboard with filtering
- ✅ Summary widgets showing payment statistics
- ✅ Responsive design
- ✅ Clean, modular component architecture

## Setup

### Prerequisites

- Node.js v20+
- npm or pnpm

### Installation

```bash
# Install dependencies
pnpm install

# Or with npm
npm install
```

### Development

```bash
# Start dev server (runs on http://localhost:5173)
pnpm dev

# Or with npm
npm run dev
```

### Build for Production

```bash
# Build the project
pnpm build

# Or with npm
npm run build

# Preview the production build
pnpm preview
```

## Configuration

The API proxy is configured in `vite.config.js` to forward requests to `http://localhost:8080/dashboard/v1`.

To use a different API endpoint, modify the proxy configuration:

```javascript
// vite.config.js
proxy: {
  '/dashboard/v1': {
    target: 'YOUR_API_URL',
    changeOrigin: true,
  }
}
```

## Authentication Flow

1. User navigates to `/login`
2. Enters email and password
3. Frontend sends POST request to `/dashboard/v1/auth/login`
4. Server returns JWT token and user role
5. Token is stored in localStorage
6. User is redirected to `/dashboard`
7. All subsequent API requests include the token in the Authorization header

## API Endpoints

### Authentication
- `POST /dashboard/v1/auth/login`
  - Body: `{ email, password }`
  - Returns: `{ token, role }`

### Payments
- `GET /dashboard/v1/payments`
  - Headers: `Authorization: Bearer {token}`
  - Returns: `{ payments: [...] }`

- `GET /dashboard/v1/payments?status=<status>`
  - Headers: `Authorization: Bearer {token}`
  - Params: `status` - one of `completed`, `processing`, `failed`
  - Returns: `{ payments: [...] }`

## Linting

```bash
pnpm lint
```

## Notes

- Tokens are stored in localStorage for persistence across page reloads
- The app uses route guards to protect the dashboard route
- API calls automatically include the bearer token from the auth store
- All dates are formatted to a readable format (e.g., "Jan 15, 2025")
- Currency amounts are formatted as USD by default
