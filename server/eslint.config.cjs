const js = require("@eslint/js");
const globals = require("globals");

/** @type {import("eslint").FlatConfig[]} */
module.exports = [
  js.configs.recommended,
  {
    languageOptions: {
      globals: {
        ...globals.browser,
      },
    },
  },
];
