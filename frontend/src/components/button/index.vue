<script setup lang="ts">

type Props = {
  /**
  ボタンの色
   */
  color: 'primary' | 'black'
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
  キャンセルボタンの有無
   */
  cancelBtn?: boolean
  /**
  キャンセルボタンのテキスト
   */
  cancelContent?: string
}

const props = withDefaults(defineProps<Props>(), {
  icon: undefined,
  color: 'primary',
  content: 'click',
  disabled: false,
  size: 'small',
  cancelBtn: false,
  cancelContent: 'cancel'
})

const emit = defineEmits<{
  (e: 'click'): void
  (e: 'cancel'): void
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
      class="margin"
      :color="color"
      :size="size"
      :variant="variant"
      :elevation="elevation"
      @click="emit('click')"
    >{{ content }}</v-btn>
    <v-btn
      v-if="cancelBtn"
      class="margin"
      :color="cancelColor"
      :size="size"
      :variant="outlined"
      :elevation="elevation"
      @click="emit('cancel')"
    >{{ cancelContent }}</v-btn>
  </div>
</template>

<style lang="scss" scoped>
.position {
  display: flex;
  justify-content: space-around;
  
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