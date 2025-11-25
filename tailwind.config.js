export default {
  content: [
    "./layouts/**/*.html",
    "./content/**/*.md",
    "./static/**/*.html",
    "./hugo.toml",
  ],
  darkMode: 'class',
  theme: {
    extend: {
      colors: {
        pixel: {
          // Цвета для совместимости и поддержки пакета
          // Все цвета определены в пакете @bavix/babichev-design
          bg: '#000000',
          'bg-dark': '#0a0a0a',
          border: '#22c55e',
          'border-light': '#4ade80',
          text: '#22c55e',
          'text-muted': '#16a34a',
          accent: '#eab308',
          'accent-dark': '#ca8a04',
          'bg-dark-dark': '#0a0a0a',
          'bg-dark-light': '#1a1a1a',
          'border-dark': '#22c55e',
          'border-dark-light': '#4ade80',
          'text-dark': '#22c55e',
          'text-dark-muted': '#16a34a',
          'accent-dark-dark': '#ca8a04',
          'bg-light': '#fafafa',
          'bg-light-dark': '#f0f0f0',
          'bg-light-light': '#ffffff',
          'border-light': '#065f46',
          'border-light-light': '#047857',
          'text-light': '#064e3b',
          'text-light-muted': '#065f46',
          'accent-light': '#c2410c',
          'accent-light-dark': '#9a3412',
        },
      },
      fontFamily: {
        pixel: ['Courier New', 'monospace'],
      },
      screens: {
        'xs': '475px',
      },
    },
  },
  plugins: [],
}
