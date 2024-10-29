<script setup lang="ts">
import { ref } from "vue";

import Button from '../components/Button/index.vue'

const count = ref(0)
console.error('count: '+count.value)
</script>
<template>
  <div>test 成功</div>
  <div>
    <Button
      color="primary"
      :content="'count: '+count"
      size="small"
      :secondBtn="true"
      :disabled="true"
      @click="count++"
    ></Button>
  </div>
</template>
