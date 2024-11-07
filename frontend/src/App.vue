<script setup lang="ts">
import Sidebar from "./navigation/Sidebar.vue";
</script>
<template>
  <div class="display">
    <div class="sidebar">
      <Sidebar />
    </div>
    <div class="body">
      <router-view />
    </div>
  </div>
</template>
<style lang="scss">
.display {
  height: 100vh;
  width: 100vw;
  display: flex;
  padding: 0 !important;
}
.sidebar {
  background-color: aliceblue;
  height: 100%;
  width: 200px;
}
.body {
  width: calc(100% - 200px);
}
</style>
