<script setup lang="ts">
import { computed, ref, watch } from 'vue';
import Button from '../Button/index.vue';

type Props={
  /**
   * 開閉
   */
  isDialog: boolean
  /**
   * タイトル
   */
  title: string
  /**
   * 内容
   */
  content: string
  /**
   * 閉じるボタンのテキスト
   */
  btnContent: string
  /**
   * ボタンの位置
   */
  btnPosition?: 'left' | 'right'
  /**
   * セカンドボタンの有無
   */
  secondBtn?: boolean
  /**
   * セカンドボタンのテキスト
   */
  secondBtnContent?: string
  /**
   * 大きさ
   */
  size?: 'small' |'medium' | 'large'
  /**
   * 永続表示
   */
  persistent?: boolean
}

const props = withDefaults(defineProps<Props>(), {
  isDialog: false,
  title: '',
  content: '',
  btnContent: '閉じる',
  persistent: false,
  secondBtn: false,
  size: 'medium'
})
const isDialog = computed(() => {
  return props.isDialog
})

const dialogSize = 'dialog-' + props.size

const emit = defineEmits<{
  (e: 'firstClick'): void
  (e: 'secondClick'): void
}>()

</script>
<template>
  <v-dialog
    v-model="isDialog"
    :persistent="persistent"
    class="dialog"
  >   
    <div class="dialog-content" :class="dialogSize">
      <div class="title">{{ title }}</div>
      <div class="content">{{ content }}</div>
      <Button
        attach="#parent"
        color="white"
        :content="btnContent"
        :position="btnPosition"
        :second-btn="secondBtn"
        :second-btn-content="secondBtnContent"
        @first-click="emit('firstClick')"
        @second-click="emit('secondClick')"
      />
    </div>
  </v-dialog>
</template>
<style lang="scss" scoped>
.dialog {
  display: flex;
  text-align: center;
}

.dialog-content {
  display: flex;
  flex-direction: column;
  margin: auto;
  background-color: white;
  text-align: start;
  padding: 14px;
  border: 1px solid black;
}

.dialog-small {
  width: 240px;
  height: 120px;
}

.dialog-medium {
  width: 450px;
  height: 240px;
}

.dialog-large {
  width: 660px;
  height: 360px;
}

.title {
  font-size: 32px;
}

.content {
  font-size: 18px;
  height: 100%;
}

</style>