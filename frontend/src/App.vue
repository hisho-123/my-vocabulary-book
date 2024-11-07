<script setup lang="ts">
import Sidebar from "@/navigation/Sidebar.vue";
import Header from "@/navigation/Header.vue";
</script>
<template>
  <div class="display">
    <div class="sidebar">
      <Sidebar />
    </div>
    <div class="page">
      <div class="header">
        <Header />
      </div>
      <div class="body">
        <router-view />
      </div>
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
.page {
  width: calc(100% - 200px);
  display: flex;
  flex-direction: column;
  padding-left: 32px;
}
.header {
  height: calc(100% / 10);
}
.sidebar {
  background-color: aliceblue;
  height: 100%;
  width: 200px;
}
</style>
