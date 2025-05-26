import type { RouteRecordRaw } from 'vue-router';

import { $t } from '#/locales';

const routes: RouteRecordRaw[] = [
  {
    name: 'About',
    path: '/about',
    component: () => import('#/views/about/index.vue'),
    meta: {
      icon: 'lucide:copyright',
      title: $t('about.title'),
      order: 9999,
    },
  },
];

export default routes;
