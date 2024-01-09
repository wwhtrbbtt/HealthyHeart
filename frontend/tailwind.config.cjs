/** @type {import('tailwindcss').Config} */
// eslint-disable-next-line no-undef
module.exports = {
  content: ["./index.html", "./src/**/*.{vue,js,ts,jsx,tsx}"],
  darkMode: ["class", '[data-theme="dark"]'],
  theme: {
    extend: {
      colors: {
        "c-gray": "#b4b4b4",
        "c-m-gray": "#495162",
        "c-m-gray-2": "#50565B",
        "c-dark-gray": "#353535",
        "button-blue": "#4154FF",
      },
    },
  },
  plugins: [],
};
