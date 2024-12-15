<script setup lang="ts">
import { ref } from "vue";
import { test } from "@/term/page.ts"

console.log("Test");
console.error("test");

const count = ref(0);
</script>
<template>
  <div>{{ test.test }}</div>
  <v-btn @click="count++">{{ count }} </v-btn>
</template>
