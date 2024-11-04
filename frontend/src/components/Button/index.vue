<script setup lang="ts">

type Props = {
  /**
  ボタンの色
   */
  color: 'primary' | 'white'
  /**
  ボタンテキスト
   */
  content: string
  /**
  活性/非活性
   */
  disabled?: boolean
  /**
  ボタンのサイズ
   */
  size?: 'small' | 'medium' | 'large'
  /**
  ボタンの位置
   */
  position?: 'left' | 'right'
  /**
  セカンドボタンの有無
   */
  secondBtn?: boolean
  /**
  セカンドボタンのテキスト
   */
  secondBtnContent?: string
}

const props = withDefaults(defineProps<Props>(), {
  icon: undefined,
  color: 'primary',
  content: 'click',
  disabled: false,
  size: 'small',
  secondBtn: false,
  secondBtnContent: 'cancel'
})

const emit = defineEmits<{
  (e: 'firstClick'): void
  (e: 'secondClick'): void
}>()

const position = 'position' + props.position

const firstBtnColor = props.secondBtn ? 'primary' : props.color === 'primary' ? 'primary' : 'black'
const variant = firstBtnColor === 'primary' ? 'elevated' : 'outlined'
const elevation = firstBtnColor === 'primary' ? 4 : 0

</script>

<template>
  <div :class="position " class="flame">
    <v-btn
      v-if="secondBtn"
      class="margin"
      color="black"
      variant="outlined"
      :size="size"
      :elevation="elevation"
      @click="emit('secondClick')"
    >{{ secondBtnContent }}</v-btn>
    <v-btn
      class="margin"
      :color="firstBtnColor"
      :size="size"
      :variant="variant"
      :elevation="elevation"
      :disabled="disabled"
      @click="emit('firstClick')"
    >{{ content }}</v-btn>
  </div>
</template>

<style lang="scss" scoped>
.position {
  display: flex;
  justify-content: center;
  
  &left {
    display: flex;
    justify-content: start;
  }
  &right {
    display: flex;
    justify-content: end;
  }
}

.flame {
  max-height: 30px;
  margin: 16px;
}

.margin {
  margin: 4px;
}

</style>