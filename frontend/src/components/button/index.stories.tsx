import { Meta, StoryObj } from "@storybook/vue3";
import Button from "./index.vue";

const meta = {
  title: 'Base/Button',
  component: Button,
  tags: ['autodocs'],
  args: {
    color: 'primary',
    content: 'BaseButton'
  },
} satisfies Meta<typeof Button>

export default meta
type Story = StoryObj<typeof meta>

export const Default: Story = {
  args: {
    color: 'black',
    content: 'BaseButton'
  }
}