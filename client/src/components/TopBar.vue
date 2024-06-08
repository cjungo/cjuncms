<template>
  <div class="top-bar">
    <ElMenu class="top-bar-menu" mode="horizontal" :default-active="appStore.topBar.defaultActive">
      <ElMenuItem index="0">
        <img alt="logo" src="/vite.svg" style="height: 90%" />
      </ElMenuItem>
      <div style="flex-grow: 1"></div>
      <ElMenuItem index="1">
        <img alt="logo" src="/vite.svg" style="height: 90%" />
      </ElMenuItem>
      
    </ElMenu>
  </div>
</template>

<script lang="ts" setup>
import { reactive } from "vue";
import { useAppStore } from "../stores/AppStore";

export type TopBarItem = {
  title: string;
  permissions?: [string];
};

const appStore = useAppStore();

const items = reactive<[TopBarItem]>([
  {
    title: "首页",
  },
]);
</script>

<style lang="scss" scoped>
.top-bar {
  display: flex;
  flex-direction: row;
  width: 100%;
}
.top-bar-menu{
  width: 100%;
}
</style>
