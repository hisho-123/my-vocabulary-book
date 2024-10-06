<script setup lang="ts">
import { ref } from 'vue';

console.log('Test')

const count = ref(0)
</script>
<template>
  <div>test 成功</div>
  <v-btn
    @click="count++"
  >{{ count }}
  </v-btn>
</template>