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
      //     title: $t('consultation.sent.title'),
      //     keepAlive: true,
      //   },
      //   name: 'SESSent',
      //   path: '/consultation/sent',
      //   component: () => import('#/views/consultation/console.vue'),
      // },
      // {
      //   meta: {
      //     title: $t('consultation.template.title'),
      //   },
      //   name: 'SESTemplate',
      //   path: '/consultation/template',
      //   component: () => import('#/views/consultation/template.vue'),
      // },
      // {
      //   meta: {
      //     title: $t('consultation.blacklist.title'),
      //     keepAlive: true,
      //   },
      //   name: 'SESBlacklist',
      //   path: '/consultation/blacklist',
      //   component: () => import('#/views/consultation/blacklist.vue'),
      // },
    ],
  },
];

export default routes;
