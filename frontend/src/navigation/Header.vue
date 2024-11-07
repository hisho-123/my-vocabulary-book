<script setup lang="ts">
import { computed } from "vue";
import { useRoute } from "vue-router";

const route = useRoute();
const title = computed(() => {
  return route.path.substring(1);
});
</script>
<template>
  <div class="header">{{ title }}</div>
</template>
<style lang="scss" scoped>
.header {
  height: 100%;
  font-size: 36px;
}
</style>
