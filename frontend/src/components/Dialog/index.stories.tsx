import { Meta, StoryObj } from "@storybook/vue3";
import Dialog from './index.vue'

const meta = {
  title: 'Base/Dialog',
  component: Dialog,
  tags: ['autodocs'],
  args: {
    isDialog: true,
    title: 'BaseDialog',
    content: 'Dialog had opened!',
    btnContent: 'close'
  },
} satisfies Meta<typeof Dialog>

export default meta
type Story = StoryObj<typeof meta>

export const Default: Story = {
  args: {
    isDialog: true,
    title: 'BaseDialog',
    content: 'Dialog had opened!',
    btnContent: 'close'
  }
}