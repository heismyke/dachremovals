export default {
  content: [
    './components/**/*.{vue,js,ts}',
    './layouts/**/*.vue',
    './pages/**/*.vue',
    './app.vue',
  ],
  theme: {
    extend: {
      colors: {
        dach: {
          orange: '#ff4f1f',
          black: '#151515',
          cream: '#fbf5f1',
          muted: '#6f6a67',
          line: '#eadfd8',
        },
      },
      fontFamily: {
        display: ['Impact', 'Arial Narrow', 'Arial', 'sans-serif'],
        sans: ['Inter', 'Avenir', 'system-ui', 'sans-serif'],
      },
    },
  },
}
