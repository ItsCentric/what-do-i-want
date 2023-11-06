/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./src/**/*.{html,go}"],
  theme: {
    extend: {
        fontFamily: {
            "chivo": ["Chivo", "ui-sans-serif"],
            "karma": ["Karma", "ui-serif"],
            "nunito": ["Nunito", "ui-sans-serif"],
        },
    },
  },
  plugins: [
      require('@tailwindcss/forms'),
  ],
}

