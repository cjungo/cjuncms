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
          <ElIcon>
            <component v-if="item.icon" :is="item.icon" />
            <IEpMenu v-else />
          </ElIcon>
          <span>{{ item.title }}</span>
        </ElMenuItem>
        <ElSubMenu :index="`${i}`" v-else>
          <template #title>
            <ElIcon>
              <component v-if="item.icon" :is="item.icon" />
              <IEpMenu v-else />
            </ElIcon>
            <span>{{ item.title }}</span>
          </template>
          <ElMenuItem
            v-for="(subItem, j) in item.children"
            :index="item.path || `${i}-${j}`"
            @click="onItemClick(subItem)"
          >
            <ElIcon>
              <component v-if="subItem.icon" :is="subItem.icon" />
              <IEpMenu v-else />
            </ElIcon>
            <span>{{ subItem.title }}</span>
          </ElMenuItem>
        </ElSubMenu>
      </template>
    </ElMenu>
  </div>
</template>

<script lang="ts" setup>
import { type Component, reactive, shallowRef } from "vue";
import { isEmpty } from "lodash";
import { useAppStore } from "../stores/AppStore";
import {
  Cpu,
  Box,
  Setting,
  Star,
  Files,
  User,
  Key,
} from "@element-plus/icons-vue";
import router from "../router";

const iconSetting = shallowRef(Setting);
const iconCpu = shallowRef(Cpu);
const iconBox = shallowRef(Box);
const iconStar = shallowRef(Star);
const iconFiles = shallowRef(Files);
const iconUser = shallowRef(User);
const iconKey = shallowRef(Key);

export type LeftBarItem = {
  title: string;
  path?: string;
  icon?: Component;
  permissions?: [string];
  children?: Array<LeftBarItem>;
};

const appStore = useAppStore();

const items = reactive<Array<LeftBarItem>>([
  {
    title: "员工管理",
    icon: iconUser,
    children: [
      {
        title: "员工列表",
        icon: iconUser,
        path: "/employee/index",
      },
    ],
  },
  {
    title: "项目管理",
    icon: iconBox,
    children: [
      {
        title: "项目列表",
        icon: iconFiles,
        path: "/project/index",
      },
      {
        title: "密钥列表",
        icon: iconKey,
        path: "/project/pass",
      },
      {
        title: "命令行",
        icon: iconCpu,
        path: "/project/shell",
      },
    ],
  },
  {
    title: "设置",
    icon: iconSetting,
    children: [
      {
        title: "我的",
        icon: iconStar,
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
  if (item.path) {
    router.push(item.path);
  }
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
