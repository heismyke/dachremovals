export default defineNuxtConfig({
  compatibilityDate: '2026-09-07',
  modules: ['@nuxtjs/tailwindcss'],
  css: ['~/assets/css/main.css'],
  experimental: {
    appManifest: false,
  },
  vite: {
    server: {
      hmr: {
        port: 24679,
      },
    },
  },
  app: {
    head: {
      title: 'Dach Removals - Moving Made Simple',
      meta: [
        {
          name: 'description',
          content:
            'Professional man and van removals across the UK. Get an instant quote for house removals, office moves, furniture delivery, and storage.',
        },
      ],
    },
  },
})
