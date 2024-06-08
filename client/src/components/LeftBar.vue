<template>
  <div class="left-bar">
    <ElMenu
      class="left-bar-menu"
      :default-active="appStore.leftBar.defaultActive"
      :collapse="appStore.leftBar.isCollapse"
      @open="onMenuOpen"
      @close="onMenuClose"
    >
      <template v-for="(item, i) in items">
        <ElMenuItem
          v-if="isEmpty(item.children)"
          :index="`${i}`"
          @click="onItemClick(item)"
        >
          <ElIcon><IEpMenu /></ElIcon>
          <span>{{ item.title }}</span>
        </ElMenuItem>
        <ElSubMenu :index="`${i}`" v-else>
          <template #title>
            <ElIcon><IEpMenu /></ElIcon>
            <span>{{ item.title }}</span>
          </template>
          <ElMenuItem
            v-for="(subItem, j) in item.children"
            :index="`${i}-${j}`"
            @click="onItemClick(subItem)"
          >
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
  path?: string;
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
        path: "/employee/index",
      },
    ],
  },
  {
    title: "项目管理",
    children: [
      {
        title: "项目列表",
        path: "/project/index",
      },
      {
        title: "命令行",
        path: "/setting/shell",
      },
    ],
  },
  {
    title: "设置",
    children: [
      {
        title: "我的",
        path: "/setting/profile",
      },
    ],
  },
]);

const onMenuOpen = (e: any) => {
  console.log("onMenuOpen", e);
};

const onMenuClose = (e: any) => {
  console.log("onMenuClose", e);
};

const onItemClick = (item: LeftBarItem) => {
  console.log("onItemClick", item);
};
</script>

<style lang="scss" scoped>
.left-bar {
  // width: 200px;
  // border-right: 1px solid #4444;

  --bg-color: #4444;
}

.left-bar-menu {
  height: 100%;
}
</style>
