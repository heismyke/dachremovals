# dachremovals

# Dach Removals Implementation

This workspace contains the three project repos for Dach Removals.

```text
/Users/myke/dachremovals
├── landing/  # Nuxt 3 public website served at https://dachremovals.co.uk
├── admin/    # Nuxt 3 admin app served at https://dachremovals.co.uk/admin
└── backend/  # Go API for auth, quotes, bookings, messages, and editable content
```

## Routing / Deployment Shape

- `https://dachremovals.co.uk` should serve the `landing` Nuxt app.
- `https://dachremovals.co.uk/admin` should serve the `admin` Nuxt app.
- The admin app is configured with `app.baseURL = '/admin/'`.
- The Go backend should run behind an API prefix such as `https://dachremovals.co.uk/api`.

Typical reverse proxy rules:

```text
/admin/* -> admin Nuxt app
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

Admin:

```bash
cd admin
npm install
npm run dev
```

Backend:

```bash
cd backend
go run ./cmd/api
```
