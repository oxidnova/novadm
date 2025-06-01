<script lang="ts" setup>
import type { ConsultationApi } from '#/api';

import { ref } from 'vue';

import { ColPage, useVbenDrawer } from '@vben/common-ui';

import { Alert, Tag } from 'ant-design-vue';

import { $t } from '#/locales';
import { formatDateFromRFC3339 } from '#/utils';

type RowType = ConsultationApi.Consultation;

const row = ref<RowType | undefined>(undefined);

const [Drawer, drawerApi] = useVbenDrawer({
  appendToMain: true,
  onClosed() {
    row.value = undefined;
  },
  onOpened() {
    row.value = drawerApi.getData<RowType>();
  },
});
</script>

<template>
  <Drawer class="w-full">
    <template #title>
      <p class="text-lg">
        {{ $t('consultation.list.moreDetailsTitle') }}
      </p>
    </template>
    <ColPage
      auto-content-height
      v-bind="{
        leftWidth: 35,
        leftMinWidth: 20,
        resizable: true,
        rightCollapsedWidth: 20,
        rightCollapsible: true,
        rightWidth: 65,
        rightMinWidth: 20,
        splitHandle: false,
        splitLine: false,
      }"
    >
      <template #title>
        <Alert
          show-icon
          :type="
            row?.status === 2
              ? 'info'
              : row?.status === 3
                ? 'success'
                : 'warning'
          "
        >
          <template #description>
            <p class="mb-2">
              <!-- <Tag class="bg-green-500">{{ $t('consultation.status') }}</Tag> -->
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
      </template>
    </ColPage>
  </Drawer>
</template>
