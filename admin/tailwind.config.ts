export default {
  content: ['./components/**/*.{vue,js,ts}', './layouts/**/*.vue', './pages/**/*.vue', './app.vue'],
  theme: {
    extend: {
      colors: {
        dach: {
          orange: '#ff4f1f',
          black: '#121212',
          cream: '#fbf5f1',
          muted: '#78716c',
          line: '#eadfd8',
        },
      },
      fontFamily: {
        display: ['Impact', 'Arial Narrow', 'Arial', 'sans-serif'],
      },
    },
  },
}
