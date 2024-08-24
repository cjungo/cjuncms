<template>
  <div class="top-bar">
    <div class="top-bar-left">
      <SearchBar />
    </div>
    <ElMenu
      class="top-bar-menu"
      mode="horizontal"
      background-color="transparent"
      :ellipsis="false"
      :default-active="appStore.topBar.defaultActive"
    >
      <div style="flex-grow: 1"></div>
      <ElMenuItem
        v-for="(item, i) in items"
        :index="`${i + 1}`"
        @click="onClickItem(item)"
      >
        <ElIcon v-if="item.icon">
          <component :is="item.icon" />
        </ElIcon>
      </ElMenuItem>
      <ElMenuItem :index="`${items.length + 1}`">
        <UserBar />
      </ElMenuItem>
    </ElMenu>
  </div>
</template>

<script lang="ts" setup>
import { type Component, ref, shallowRef } from "vue";
import { useAppStore } from "../stores/AppStore";
import { HomeFilled } from "@element-plus/icons-vue";
import { useRouter } from "vue-router";

type TopBarItem = {
  title: string;
  icon?: Component;
  path?: string;
  permissions?: [string];
};
const home = shallowRef(HomeFilled);
const appStore = useAppStore();
const router = useRouter();

const items = ref<[TopBarItem]>([
  {
    title: "首页",
    icon: home,
    path: "/index",
  },
]);

const onClickItem = (item: TopBarItem) => {
  if (item.path) {
    router.push(item.path);
  }
};
</script>

<style lang="scss" scoped>
.top-bar {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  align-items: stretch;
  width: 100%;
  box-sizing: border-box;
  // background-color: #fbfaff;
}
.top-bar-left {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding-left: 1em;
  border-bottom: 1px solid var(--el-menu-border-color);
}
.top-bar-menu {
  flex-grow: 1;
}
</style>
