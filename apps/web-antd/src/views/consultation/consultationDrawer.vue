<script lang="ts" setup>
import type { ConsultationApi } from '#/api';

import { ref } from 'vue';

import { useVbenDrawer } from '@vben/common-ui';

import { Alert, Button, Card, Tag } from 'ant-design-vue';

import { $t } from '#/locales';
import { formatDateFromRFC3339 } from '#/utils';

import MarkdownEditor from '../components/markdownEditor.vue';

type RowType = ConsultationApi.Consultation;

const row = ref<RowType | undefined>(undefined);

const [Drawer, drawerApi] = useVbenDrawer({
  appendToMain: true,
  onClosed() {
    row.value = undefined;
  },
  onOpened() {
    row.value = drawerApi.getData<RowType>();
    mdContent.value = row.value?.content || '';
  },
});

const generateImage = () => {};

const mdContent = ref('');
</script>

<template>
  <Drawer class="w-full">
    <template #title>
      <p class="text-lg">
        {{ $t('consultation.list.moreDetailsTitle') }}
      </p>
    </template>

    <Alert
      show-icon
      :type="
        row?.status === 2 ? 'info' : row?.status === 3 ? 'success' : 'warning'
      "
    >
      <template #description>
        <p class="mb-2">
          <Tag
            class="bg-green-500"
            :color="
              row?.status === 2
                ? '#cca43f'
                : row?.status === 3
                  ? '#87d068'
                  : '#f50'
            "
          >
            {{
              row?.status === 2
                ? 'Draft'
                : row?.status === 3
                  ? 'Publiched'
                  : 'Fetching'
            }}
          </Tag>
          <Tag class="bg-green-500">
            {{ $t('consultation.createdAt') }}
          </Tag>
          {{ formatDateFromRFC3339(row?.createdAt) }}
          <Tag class="bg-green-500">ID</Tag>
          {{ row?.id }}
          <Tag class="bg-green-500">{{ $t('consultation.prompt') }}</Tag>
          {{ row?.prompt }}
        </p>
      </template>
    </Alert>

    <Card class="mt-2" :body-style="{ padding: 0 }">
      <template #title>
        <div class="flex items-center justify-between gap-2">
          <span>{{ $t('consultation.content') }}</span>
          <Button type="primary" size="small" danger @click="generateImage()">
            {{ $t('consultation.list.exportImage') }}
          </Button>
        </div>
      </template>
      <div
        class="relative flex min-h-32 items-center justify-center gap-2 overflow-hidden"
      >
        <MarkdownEditor
          style="height: calc(100vh - 320px)"
          v-model="mdContent"
        />
      </div>
    </Card>
  </Drawer>
</template>
