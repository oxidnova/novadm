<script lang="ts" setup>
import { MdEditor } from 'md-editor-v3';

import 'md-editor-v3/lib/style.css';
import '@vavt/cm-extension/dist/previewTheme/arknights.css';

// 定义组件 props
const props = defineProps({
  modelValue: {
    type: String,
    required: true,
  },
});

// 定义 emits
const emit = defineEmits(['update:modelValue']);

// 处理编辑器内容变化
const handleChange = (value: string) => {
  emit('update:modelValue', value);
};
</script>

<template>
  <MdEditor
    preview-theme="arknights"
    :model-value="props.modelValue"
    @on-change="handleChange"
    :toolbars-exclude="['save', 'htmlPreview', 'github']"
  />
</template>

<style>
.arknights-theme ul {
  list-style-type: disc;
}
</style>
