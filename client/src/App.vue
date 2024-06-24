<template>
  <ElConfigProvider :locale="EL_LOCALE" :size="EL_SIZE" :z-index="EL_ZINDEX">
    <RouterView></RouterView>
  </ElConfigProvider>
</template>

<script setup lang="ts">
import { ElConfigProvider } from "element-plus";
import zhCn from "element-plus/dist/locale/zh-cn.mjs";
import { onBeforeMount } from "vue";
import { delay } from "./utils/time";
import { useAuthStore } from "./stores/AuthStore";
import { isEmpty } from "lodash";
import { renewal } from "./apis/login";

const EL_LOCALE = zhCn;
const EL_SIZE = "small";
const EL_ZINDEX = 4000;

onBeforeMount(async () => {
  const auth = useAuthStore();

  while (true) {
    await delay(60 * 60 * 1000);
    if (!isEmpty(auth.token)) {
      const result = await renewal();
      auth.token = result.data;
    }
  }
});
</script>

<style lang="scss" scoped></style>
