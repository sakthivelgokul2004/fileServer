/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {},
    colors: {
      'blue': '#1fb6ff',
      'purple': '#7e5bef',
      'pink': '#ff49db',
      'orange': '#ff7849',
      'bg': '#0f172a',
      'boder': '#1d2537',
      'text-hed': '#E2E8F0',
      'green': '#13ce66',
      'yellow': '#ffc82c',
      'gray-dark': '#273444',
      'gray': '#8492a6',
      'gray-light': '#d3dce6',
      'eample':'#1e293b',
      'inputbg':'#142942',
      'link':'#0ea5e9'
    },
  },
  plugins: [require('daisyui'),],
}
