module.exports = {
  root: true,
  env: {
    browser: true,
    node: true
  },
  parserOptions: { parser: "babel-eslint" },
  extends: ["prettier", "@nuxtjs", "plugin:nuxt/recommended"],
  // add your custom rules here
  rules: {}
};
