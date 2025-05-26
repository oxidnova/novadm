<script lang="ts" setup>
import { Page } from '@vben/common-ui';

import { $t } from '#/locales';

defineOptions({ name: 'About' });
</script>

<template>
  <Page :title="$t('about.pageTitle')">
    <template #description>
      <p class="text-foreground mt-3 text-sm leading-6">
        <a class="vben-link" target="_blank"> SysForge </a>
        {{ $t('about.description') }}
      </p>
    </template>
  </Page>
</template>
