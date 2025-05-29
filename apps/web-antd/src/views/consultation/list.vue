<script lang="ts" setup>
import type { VbenFormProps } from '#/adapter/form';
import type { VxeGridProps } from '#/adapter/vxe-table';
import type { ConsultationApi } from '#/api';

import { Page } from '@vben/common-ui';

import { message } from 'ant-design-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { $t } from '#/locales';

type RowType = ConsultationApi.Consultation;

const formOptions: VbenFormProps = {
  // 默认展开
  collapsed: false,
  fieldMappingTime: [['date', ['start', 'end']]],
  schema: [
    // {
    //   component: markRaw(QueryEmailField),
    //   defaultValue: ['to', ''],
    //   disabledOnChangeListener: false,
    //   fieldName: 'queryEmail',
    //   formItemClass: 'col-span-1',
    //   label: $t('consultation.email'),
    // },
    {
      component: 'Select',
      componentProps: {
        allowClear: true,
        filterOption: true,
        options: [
          {
            label: 'ERROR',
            value: 'SendErr',
          },
          {
            label: 'SendOK',
            value: 'SendOK',
          },
        ],
      },
      fieldName: 'status',
      label: $t('consultation.status'),
    },
    {
      component: 'Input',
      fieldName: 'eid',
      label: 'EID',
    },
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
      field: 'email.from',
      slots: { default: 'from' },
      title: 'From',
      width: 280,
    },
    {
      field: 'email.to',
      slots: { default: 'to' },
      title: 'To',
      width: 280,
    },
    {
      field: 'status',
      slots: { default: 'status' },
      title: $t('consultation.status'),
      width: 100,
    },
    {
      field: 'platform',
      title: $t('consultation.platform'),
      width: 100,
    },
    {
      field: 'send_date',
      slots: { default: 'sendDate' },
      title: $t('consultation.sendAt'),
      width: 220,
    },
    {
      field: 'eid',
      slots: { default: 'eid' },
      title: 'EID',
      minWidth: 280,
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
</script>

<template>
  <Page auto-content-height>
    <Grid>
      <div>consultation list</div>
    </Grid>
  </Page>
</template>
