/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ['./pkg/views/**/*.templ', './pkg/views/*.templ'],
  theme: {
    extend: {
      fontFamily: {
        teko: ['Teko', 'sans-serif'],
      },
    },
  },
  plugins: [require('@tailwindcss/typography'), require('daisyui')],
  daisyui: {
    themes: ['forest'],
    base: true,
    styled: true,
    utils: true,
  },
};
