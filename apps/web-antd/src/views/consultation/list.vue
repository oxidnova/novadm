<script lang="ts" setup>
import type { VbenFormProps } from '#/adapter/form';
import type { VxeGridProps } from '#/adapter/vxe-table';
import type { ConsultationApi } from '#/api';

import { Page } from '@vben/common-ui';

import { Button, message, Tag } from 'ant-design-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { $t } from '#/locales';
import { formatDateFromTimestamp } from '#/utils';

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
            label: 'Publiched',
            value: 1,
          },
          {
            label: 'UnPubliched',
            value: 0,
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
      width: 220,
    },
    {
      field: 'prompt',
      slots: { default: 'prompt' },
      title: $t('consultation.prompt'),
      width: 180,
    },
    {
      field: 'status',
      slots: { default: 'status' },
      title: $t('consultation.status'),
      width: 100,
    },
    {
      field: 'content',
      slots: { default: 'content' },
      title: $t('consultation.content'),
      minWidth: 220,
    },
    {
      field: 'updatedAt',
      slots: { default: 'updatedAt' },
      title: $t('consultation.updatedAt'),
      width: 220,
    },
    {
      field: 'id',
      slots: { default: 'id' },
      title: 'ID',
      width: 200,
    },
    {
      field: 'action',
      slots: { default: 'action' },
      fixed: 'right',
      title: $t('page.action'),
      width: 120,
    },
  ],
  exportConfig: {},
  height: 'auto',
  keepSource: true,
  pagerConfig: {},
  proxyConfig: {
    ajax: {
      query: async ({ page }, formValues) => {
        const params = {
          ...formValues,
          page: page.currentPage,
          limit: page.pageSize,
        };
        // const params = toFetchParams(
        //   formValues,
        //   page.currentPage,
        //   page.pageSize,
        // );
        message.success(`Query params: ${JSON.stringify(params)}`);
        // const data = await searchEmailsApi(params);
        // pageCache.set(page.currentPage, data.lastId);
        // return data;
        return {};
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

const openDrawer = (_row: RowType) => {
  // drawerApi.setState({ placement: 'left' }).setData<RowType>(row).open();
};
</script>

<template>
  <Page auto-content-height>
    <Grid>
      <template #createdAt="{ row }">
        <Tag>{{ formatDateFromTimestamp(row.createdAt) }}</Tag>
      </template>
      <template #prompt="{ row }">
        <Tag>{{ row.prompt }}</Tag>
      </template>
      <template #status="{ row }">
        <Tag :color="row.status === 1 ? '#87d068' : '#f50'">
          {{ row.status === 1 ? 'Published' : 'UnPublished' }}
        </Tag>
      </template>
      <template #content="{ row }">
        {{ row.content }}
      </template>
      <template #updatedAt="{ row }">
        <Tag>{{ formatDateFromTimestamp(row.updatedAt) }}</Tag>
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
  </Page>
</template>
