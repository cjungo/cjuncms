<template>
  <div class="left-bar">
    <ElMenu
      :default-active="appStore.leftBar.defaultActive"
      :collapse="appStore.leftBar.isCollapse"
      @open="onMenuOpen"
      @close="onMenuClose"
    >
      <template v-for="item in items">
        <ElMenuItem v-if="isEmpty(item.children)">
          <ElIcon><IEpMenu /></ElIcon>
          <span>{{ item.title }}</span>
        </ElMenuItem>
        <ElSubMenu v-else>
          <template #title>
            <ElIcon><IEpMenu /></ElIcon>
            <spen>{{ item.title }}</spen>
          </template>
          <ElMenuItem v-for="subItem in item.children">
            <ElIcon><IEpMenu /></ElIcon>
            <span>{{ subItem.title }}</span>
          </ElMenuItem>
        </ElSubMenu>
      </template>
    </ElMenu>
  </div>
</template>

<script lang="ts" setup>
import { reactive } from "vue";
import { isEmpty } from "lodash";
import { useAppStore } from "../stores/AppStore";

export type LeftBarItem = {
  title: string;
  permissions?: [string];
  children?: Array<LeftBarItem>;
};

const appStore = useAppStore();

const items = reactive<Array<LeftBarItem>>([
  {
    title: "员工管理",
    children: [
      {
        title: "员工列表",
      },
    ],
  },
  {
    title: "查看",
  },
]);

const onMenuOpen = (e: any) => {
  console.log("onMenuOpen", e);
};

const onMenuClose = (e: any) => {
  console.log("onMenuClose", e);
};
</script>

<style lang="scss" scoped>
.left-bar {
  // width: 200px;
  // border-right: 1px solid #4444;

  --bg-color: #4444;
}
</style>
