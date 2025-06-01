<script lang="ts" setup>
import type { VbenFormProps } from '#/adapter/form';
import type { VxeGridProps } from '#/adapter/vxe-table';
import type { ConsultationApi } from '#/api';

import { EllipsisText, Page, useVbenDrawer } from '@vben/common-ui';

import { Button, message, Tag } from 'ant-design-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { searchConsultationsApi } from '#/api';
import { $t } from '#/locales';
import { formatDateFromRFC3339 } from '#/utils';

import ConsultationDrawer from './consultationDrawer.vue';

type RowType = ConsultationApi.Consultation;

const formOptions: VbenFormProps = {
  // 默认展开
  collapsed: false,
  fieldMappingTime: [['date', ['start', 'end']]],
  schema: [
    {
      component: 'RangePicker',
      // defaultValue: [dayjs().subtract(1, 'days'), dayjs()],
      fieldName: 'dateRange',
      label: $t('consultation.dateRange'),
    },
    {
      component: 'TimePicker',
      // defaultValue: dayjs().hour(0).minute(0).second(0),
      fieldName: 'timeRangeStart',
      label: $t('consultation.startTime'),
    },
    {
      component: 'TimePicker',
      // defaultValue: '00:00:00',
      fieldName: 'timeRangeEnd',
      label: $t('consultation.endTime'),
    },
    {
      component: 'Input',
      fieldName: 'id',
      label: 'ID',
    },
    {
      component: 'Select',
      componentProps: {
        allowClear: true,
        filterOption: true,
        options: [
          {
            label: 'Draft',
            value: 2,
          },
          {
            label: 'Fetching',
            value: 1,
          },
          {
            label: 'Publiched',
            value: 3,
          },
        ],
      },
      fieldName: 'status',
      label: $t('consultation.status'),
    },
  ],
  // 控制表单是否显示折叠按钮
  showCollapseButton: true,
  // 是否在字段值改变时提交表单
  submitOnChange: false,
  // 按下回车时是否提交表单
  submitOnEnter: false,
};

const getFetchParams = (
  formValues: any,
  currentPage: number,
  pageSize: number,
): ConsultationApi.FetchParams => {
  const startTime = formValues.dateRange?.[0]
    ?.hour(formValues.timeRangeStart?.hour() ?? 0)
    .minute(formValues.timeRangeStart?.minute() ?? 0)
    .second(formValues.timeRangeStart?.second() ?? 0)
    .unix();
  const endTime = formValues?.dateRange?.[1]
    ?.hour(formValues.timeRangeEnd?.hour() ?? 23)
    .minute(formValues.timeRangeEnd?.minute() ?? 59)
    .second(formValues.timeRangeEnd?.second() ?? 59)
    .unix();

  const params: ConsultationApi.FetchParams = {
    page: currentPage,
    pageSize,
    startTime,
    endTime,
    id: formValues.id,
    status: formValues.status,
  };

  return params;
};

const gridOptions: VxeGridProps<RowType> = {
  checkboxConfig: {
    highlight: true,
    labelField: 'name',
  },
  columns: [
    { title: $t('page.tab.seq'), type: 'seq', width: 50 },
    {
      field: 'createdAt',
      slots: { default: 'createdAt' },
      title: $t('consultation.createdAt'),
      width: 200,
    },
    {
      field: 'status',
      slots: { default: 'status' },
      title: $t('consultation.status'),
      width: 100,
    },
    {
      field: 'prompt',
      slots: { default: 'prompt' },
      title: $t('consultation.prompt'),
      width: 140,
    },
    {
      field: 'content',
      slots: { default: 'content' },
      title: $t('consultation.content'),
      minWidth: 400,
    },
    {
      field: 'updatedAt',
      slots: { default: 'updatedAt' },
      title: $t('consultation.updatedAt'),
      width: 200,
    },
    {
      field: 'id',
      slots: { default: 'id' },
      title: 'ID',
      width: 280,
    },
    {
      field: 'action',
      fixed: 'right',
      slots: { default: 'action' },
      title: $t('page.action'),
      width: 120,
    },
  ],
  exportConfig: {},
  height: 'auto',
  keepSource: true,
  pagerConfig: {
    pageSize: 10,
  },
  proxyConfig: {
    ajax: {
      query: async ({ page }, formValues) => {
        const params = getFetchParams(
          formValues,
          page.currentPage,
          page.pageSize,
        );
        message.success(`Query params: ${JSON.stringify(params)}`);
        const data = await searchConsultationsApi(params);
        return data;
      },
    },
  },
  toolbarConfig: {
    custom: true,
    // export: true,
    // refresh: true,
    resizable: true,
    // search: true,
    zoom: true,
  },
  showOverflow: false,
};

const [Grid] = useVbenVxeGrid({
  formOptions,
  gridOptions,
});

const [Drawer, drawerApi] = useVbenDrawer({
  connectedComponent: ConsultationDrawer,
  showConfirmButton: false,
  showCancelButton: false,
});

const openDrawer = (row: RowType) => {
  drawerApi.setState({ placement: 'left' }).setData<RowType>(row).open();
};
</script>

<template>
  <Page auto-content-height>
    <Grid>
      <template #createdAt="{ row }">
        <Tag>{{ formatDateFromRFC3339(row.createdAt) }}</Tag>
      </template>
      <template #status="{ row }">
        <Tag
          :color="
            row?.status === 2
              ? '#cca43f'
              : row?.status === 3
                ? '#87d068'
                : '#f50'
          "
        >
          {{
            row.status === 2
              ? 'Draft'
              : row.status === 3
                ? 'Publiched'
                : 'Fetching'
          }}
        </Tag>
      </template>
      <template #prompt="{ row }">
        <Tag>{{ row.prompt }}</Tag>
      </template>
      <template #content="{ row }">
        <EllipsisText :line="1" expand>
          {{ row.content }}
        </EllipsisText>
      </template>
      <template #updatedAt="{ row }">
        <Tag>{{ formatDateFromRFC3339(row.updatedAt) }}</Tag>
      </template>
      <template #id="{ row }">
        <Tag color=""> {{ row.id }}</Tag>
      </template>
      <template #action="{ row }">
        <Button type="link" @click="openDrawer(row)">
          {{ $t('page.tab.moreDetails') }}
        </Button>
      </template>
    </Grid>
    <Drawer />
  </Page>
</template>
