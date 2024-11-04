import vue from 'eslint-plugin-vue';
import prettier from 'eslint-config-prettier';
import vueParser from 'vue-eslint-parser';

export default [
  {
    files: ['**/*.{ts,tsx,vue}'],
    ignores: ['node_modules'],
    languageOptions: {
      parser: vueParser,
      parserOptions: {
        parser: '@typescript-eslint/parser',
        ecmaVersion: 'latest',
        sourceType: 'module',
      },
      ecmaVersion: 'latest',
      sourceType: 'module',
      globals: {
        browser: true,
        'vue/setup-compiler-macros': true,
      },
    },
    plugins: {
      vue: vue,
    },
    rules: {
      ...vue.configs['vue3-recommended'].rules,
      ...prettier.rules,
    },
  },
];