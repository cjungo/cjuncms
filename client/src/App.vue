<template>
  <ElConfigProvider :locale="EL_LOCALE" :size="EL_SIZE" :z-index="EL_ZINDEX">
    <RouterView></RouterView>
  </ElConfigProvider>
</template>

<script setup lang="ts">
import { ElConfigProvider } from 'element-plus';
import zhCn from 'element-plus/dist/locale/zh-cn.mjs';

const EL_LOCALE = zhCn;
const EL_SIZE = 'small';
const EL_ZINDEX = 4000;
</script>

<style lang="scss" scoped>

</style>
