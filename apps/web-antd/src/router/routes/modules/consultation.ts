import type { RouteRecordRaw } from 'vue-router';

import { $t } from '#/locales';

const routes: RouteRecordRaw[] = [
  {
    meta: {
      icon: 'emojione-monotone:newspaper',
      keepAlive: true,
      order: 10,
      title: $t('consultation.title'),
      authority: ['all', 'consultation'],
    },
    name: 'Consultation',
    path: '/consultation',
    children: [
      {
        meta: {
          title: $t('consultation.list.title'),
          keepAlive: true,
        },
        name: 'ConsultationList',
        path: '/consultation/list',
        component: () => import('#/views/consultation/list.vue'),
      },
      // {
      //   meta: {
      //     title: '测试',
      //     keepAlive: true,
      //   },
      //   name: 'consultationTest',
      //   path: '/consultation/test',
      //   component: () => import('#/views/consultation/test.vue'),
      // },
    ],
  },
];

export default routes;
