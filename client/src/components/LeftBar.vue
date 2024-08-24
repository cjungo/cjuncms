<template>
  <div class="left-bar">
    <div class="left-bar-head">
      <img alt="logo" src="/vite.svg" style="height: 50%" />
    </div>
    <ElMenu
      class="left-bar-menu"
      :default-active="appStore.leftBar.defaultActive"
      :collapse="appStore.leftBar.isCollapse"
      text-color="#fff"
      background-color="transparent"
      active-text-color="#49f"
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
            <component
              v-else-if="item.path && metas[item.path]?.icon"
              :is="metas[item.path].icon"
            />
            <IEpMenu v-else />
          </ElIcon>
          <span>{{ item.title }}</span>
        </ElMenuItem>
        <ElSubMenu :index="`${i}`" v-else>
          <template #title>
            <ElIcon>
              <component v-if="item.icon" :is="item.icon" />
              <component
                v-else-if="item.path && metas[item.path]?.icon"
                :is="metas[item.path].icon"
              />
              <IEpMenu v-else />
            </ElIcon>
            <span>{{ item.title }}</span>
          </template>
          <ElMenuItem
            v-for="(subItem, j) in item.children"
            :index="subItem.path"
            @click="onItemClick(subItem)"
          >
            <ElIcon>
              <component v-if="subItem.icon" :is="subItem.icon" />
              <component
                v-else-if="subItem.path && metas[subItem.path]?.icon"
                :is="metas[subItem.path].icon"
              />
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
import { metas } from "../router";
import { useAppStore } from "../stores/AppStore";
import { Box, Setting, User } from "@element-plus/icons-vue";
import router from "../router";

const iconSetting = shallowRef(Setting);
const iconBox = shallowRef(Box);
const iconUser = shallowRef(User);

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
        path: "/employee/index",
      },
      {
        title: "部门管理",
        path: "/employee/department",
      },
    ],
  },
  {
    title: "项目管理",
    icon: iconBox,
    children: [
      {
        title: "项目列表",
        path: "/project/index",
      },
      {
        title: "密钥列表",
        path: "/project/pass",
      },
      {
        title: "命令行",
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

  background: 100% 100% url(../assets/left-bar-bg.png);
}

.left-bar-menu {
  height: 100%;
}

.left-bar-head {
  display: flex;
  align-items: center;
  justify-content: center;

  box-sizing: border-box;
  height: var(--el-menu-horizontal-height);
  border-bottom: 1px solid var(--el-menu-border-color);
  border-right: 1px solid var(--el-menu-border-color);
}
</style>
