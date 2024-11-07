<script setup lang="ts">
import { useRoute } from "vue-router";
import type { navigateList } from "./navigation.ts";
import { applicationName } from "./navigation.ts";
import "material-symbols";
import { computed } from "vue";

const route = useRoute();
const currentUrl = computed(() => {
  return route.path.substring(1)
});

// TODO: DBから単語帳をとってきてオブジェクトの配列にする
const sidebarList: navigateList = [
  { name: "リスト１", link: "/test" },
  { name: "list2", link: "/yaka" },
];
</script>
<template>
  <div class="sidebar">
    <div class="appName">{{ applicationName }}</div>
    <p v-for="list in sidebarList">
      <router-link :to="list.link" class="sidebar-list">
        <span
          v-if="currentUrl && list.link.includes(currentUrl)"
          class="material-symbols-outlined fill"
          >Label</span
        >
        <span v-else class="material-symbols-outlined">Label</span>
        <div>{{ list.name }}</div>
      </router-link>
    </p>
  </div>
</template>
<style lang="scss" scoped>
.appName {
  height: 50px;
  font-size: 20px;
  padding: 8px;
}
.sidebar {
  width: 100%;
  height: 100%;
  text-align: left;
}
.material-symbols-outlined {
  padding-right: 4px;
}
.fill {
  font-variation-settings: "FILL" 1;
}
.sidebar-list {
  display: flex;
  padding: 6px;
}
</style>
