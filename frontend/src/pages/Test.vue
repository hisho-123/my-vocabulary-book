<script setup lang="ts">
import { ref } from "vue";

console.log("Test");
console.error("test");

const count = ref(0);
</script>
<template>
  <div>test 成功</div>
  <v-btn @click="count++">{{ count }} </v-btn>
</template>
