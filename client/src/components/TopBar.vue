<template>
  <div class="top-bar">
    <ElMenu
      class="top-bar-menu"
      mode="horizontal"
      :ellipsis="false"
      :default-active="appStore.topBar.defaultActive"
    >
      <ElMenuItem index="0">
        <div class="top-bar-head">
          <img alt="logo" src="/vite.svg" style="height: 50%" />
        </div>
      </ElMenuItem>
      <div style="flex-grow: 1"></div>
      <ElMenuItem v-for="(item, i) in items" :index="`${i + 1}`">
        {{ item.title }}
      </ElMenuItem>
      <ElMenuItem :index="`${items.length + 1}`">
        <UserBar/>
      </ElMenuItem>
    </ElMenu>
  </div>
</template>

<script lang="ts" setup>
import { reactive } from "vue";
import { useAppStore } from "../stores/AppStore";

type TopBarItem = {
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
.top-bar-menu {
  width: 100%;
}

.top-bar-head {
  display: flex;
  align-items: center;
  justify-content: center;
}
</style>
