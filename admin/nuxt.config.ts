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
        port: 24680,
      },
    },
  },
  app: {
    baseURL: '/admin/',
    head: {
      title: 'Dach Removals Admin',
    },
  },
})
