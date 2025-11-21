export default {
  content: [
    "./layouts/**/*.html",
    "./content/**/*.md",
    "./static/**/*.html",
    "./hugo.toml",
  ],
  safelist: [
    'bg-pixel-bg-dark',
    'bg-pixel-bg',
    'bg-pixel-bg-light',
    'text-pixel-text',
    'text-pixel-text-muted',
    'text-pixel-accent',
    'text-pixel-border-light',
    'border-pixel-border',
    'border-pixel-border-light',
    'pixel-box',
    'pixel-button',
    'pixel-text',
    'error-terminal-line',
  ],
  theme: {
    extend: {
      colors: {
        pixel: {
          bg: '#000000',
          'bg-dark': '#0a0a0a',
          'bg-light': '#1a1a1a',
          border: '#22c55e',
          'border-light': '#4ade80',
          text: '#22c55e',
          'text-muted': '#16a34a',
          accent: '#eab308',
          'accent-dark': '#ca8a04',
        },
      },
      fontFamily: {
        pixel: ['Courier New', 'monospace'],
      },
      animation: {
        'blink': 'blink 1s step-end infinite',
        'fade-in': 'fadeIn 0.3s ease-in',
      },
      keyframes: {
        blink: {
          '0%, 50%': { opacity: '1' },
          '51%, 100%': { opacity: '0' },
        },
        fadeIn: {
          '0%': { opacity: '0' },
          '100%': { opacity: '1' },
        },
      },
      screens: {
        'xs': '475px',
      },
    },
  },
  plugins: [],
}
