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
  secondContent?: string
}

const props = withDefaults(defineProps<Props>(), {
  icon: undefined,
  color: 'primary',
  content: 'click',
  disabled: false,
  size: 'small',
  secondBtn: false,
  secondContent: 'cancel'
})

const emit = defineEmits<{
  (e: 'click'): void
  (e: 'second'): void
}>()

// 固定値を代入
const position = 'position-' + props.position
const elevation = 4
const cancelColor = 'black'
const outlined = 'outlined'

const variant = props.color === 'primary' ? 'elevated' : 'outlined'
console.error('color: '+props.color)
console.error('variant: '+variant)
</script>

<template>
  <div :class="position " class="flame">
      <v-btn
      v-if="secondBtn"
      class="margin"
      :color="cancelColor"
      :size="size"
      :variant="outlined"
      :elevation="elevation"
      :disabled="disabled"
      @click="emit('second')"
    >{{ secondContent }}</v-btn>
    <v-btn
      class="margin"
      :color="color"
      :size="size"
      :variant="variant"
      :elevation="elevation"
      :disabled="disabled"
      @click="emit('click')"
    >{{ content }}</v-btn>
  </div>
</template>

<style lang="scss" scoped>
.position {
  display: flex;
  justify-content: center;
  
  &-left {
    display: flex;
    justify-content: start;
  }
  &-right {
    display: flex;
    justify-content: end;
  }
}

.flame {
  min-width: 700px;
  max-height: 30px;
  margin: 16px;

}

.margin {
  margin: 4px;
}

</style>