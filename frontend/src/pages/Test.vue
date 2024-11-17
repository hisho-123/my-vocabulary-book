<script setup lang="ts">
import { ref } from "vue";
import { test } from "@/term/page.ts"

import Button from '@/components/Button/index.vue'
import Dialog from '@/components/Dialog/index.vue'

const count = ref(0)
const dialog = ref(false)

</script>
<template>
  <div>{{ test.test }}</div>
  <div>
    <Button
      color="primary"
      :content="'count: '+count"
      size="small"
      :secondBtn="true"
      secondBtnContent="count=0"
      @firstClick="dialog = !dialog"
      @secondClick="count=0"
    ></Button>
  </div>
  <Dialog
    :isDialog="dialog"
    title="title"
    content="contentcontent"
    btnContent="count+1"
    secondBtnContent="cancel"
    :secondBtn="true"
    @firstClick="count++"
    @secondClick="dialog = !dialog"
  />
</template>
