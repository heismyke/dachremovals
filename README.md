# dachremovals

Dach Removals Nuxt frontend for the public website and admin route.

## Routing / Deployment Shape

- `https://dachremovals.co.uk` serves the public landing page.
- `https://dachremovals.co.uk/admin` serves the admin dashboard routes in this same Nuxt app.
- Backend/API code lives in `git@github.com:heismyke/dachremovals-engine.git`.

Typical reverse proxy rules:

```text
/*       -> landing Nuxt app
```

## Local Development

```bash
npm install
npm run dev
```
