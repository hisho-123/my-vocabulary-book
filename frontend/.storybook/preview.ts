import { setup, type Preview } from "@storybook/vue3";
import { createApp } from 'vue';
import { createVuetify } from 'vuetify';
import 'vuetify/styles';
import * as components from 'vuetify/components';
import * as directives from 'vuetify/directives';

const preview: Preview = {
  parameters: {
    controls: {
      matchers: {
        color: /(background|color)$/i,
        date: /Date$/i,
      },
    },
  },
};

export default preview;

const app = createApp({})

const vuetify=createVuetify({
  components,
  directives,
})

setup((app) => {
  app.use(vuetify);
});