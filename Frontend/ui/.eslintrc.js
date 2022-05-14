module.exports = {
  root: true,
  env: {
    node: true
  },
  extends: [
    'plugin:vue/vue3-essential',
    'eslint:recommended',
    '@vue/typescript/recommended'
  ],
  parserOptions: {
    ecmaVersion: 2020
  },
  rules: {
    "@typescript-eslint/camelcase": "off",
    // 'vue/valid-v-slot': ['error', {allowModifiers: true,}]
  // '@typescript-eslint/no-var-requires': 'off'
  }
};