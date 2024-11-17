<script setup lang="ts">
import { ref } from "vue";
import { test } from "@/term/page.ts"

import Button from '../components/button/index.vue'

const count = ref(0)
console.error('count: '+count.value)
</script>
<template>
  <div>{{ test.test }}</div>
  <div>
    <Button
      color="primary"
      :content="'count: '+count"
      size="small"
      :cancelBtn="true"
      position="right"
      @click="count++"
    ></Button>
  </div>
</template>
