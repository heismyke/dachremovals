export default defineNuxtConfig({
  compatibilityDate: '2026-09-07',
  modules: ['@nuxtjs/tailwindcss'],
  css: ['~/assets/css/main.css'],
  runtimeConfig: {
    public: {
      apiBase: process.env.NUXT_PUBLIC_API_BASE || 'http://localhost:8080',
    },
  },
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
      title: 'Dach Removals | UK Removals, Man and Van, House Moves',
      htmlAttrs: {
        lang: 'en-GB',
      },
      meta: [
        { charset: 'utf-8' },
        { name: 'viewport', content: 'width=device-width, initial-scale=1' },
        { name: 'theme-color', content: '#ff4f1f' },
        {
          name: 'description',
          content:
            'Dach Removals provides professional UK removals, man and van services, house moves, office relocations, furniture delivery, packing, and storage support.',
        },
      ],
      link: [
        { rel: 'icon', type: 'image/svg+xml', href: '/favicon.svg' },
        { rel: 'shortcut icon', href: '/favicon.svg' },
        { rel: 'apple-touch-icon', href: '/images/logo.jpg' },
        { rel: 'manifest', href: '/site.webmanifest' },
      ],
    },
  },
})
