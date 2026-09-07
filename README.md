# dachremovals

# Dach Removals Implementation

This workspace contains the Dach Removals Nuxt frontend and Go backend.

```text
/Users/myke/dachremovals
├── landing/  # Nuxt 3 public website served at https://dachremovals.co.uk
└── backend/  # Go API for auth, quotes, bookings, messages, and editable content
```

## Routing / Deployment Shape

- `https://dachremovals.co.uk` should serve the `landing` Nuxt app.
- `https://dachremovals.co.uk/admin` should serve the admin routes inside the same `landing` Nuxt app.
- The Go backend should run behind an API prefix such as `https://dachremovals.co.uk/api`.

Typical reverse proxy rules:

```text
/api/*   -> Go backend
/*       -> landing Nuxt app
```

## Local Development

Landing:

```bash
cd landing
npm install
npm run dev
```

Backend:

```bash
cd backend
go run ./cmd/api
```
