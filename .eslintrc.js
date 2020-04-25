module.exports = {
  root: true,
  env: {
    browser: true,
    node: true
  },
  extends: [
    "eslint:recommended",
    "plugin:vue/recommended"
  ],
  // add your custom rules here
  plugins: ["vue"],
  rules: {
    "no-console": 0,
    "computed-property-spacing": 2,
    "no-var": 2,
    "no-multiple-empty-lines": [2, {"max": 1}],
    "keyword-spacing": 2,
    "func-call-spacing": [2, "never"],
  }
}
