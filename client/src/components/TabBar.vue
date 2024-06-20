<template>
  <div class="tab-bar">
    <ElTabs
      v-model="appStore.tabBar.activeName"
      type="card"
      @tab-change="onTabChange"
      @tab-click="onTabClick"
      @tab-remove="onTabRemove"
    >
      <ElTabPane label="首页" name="/index" :closable="false"></ElTabPane>
      <ElTabPane
        v-for="item in appStore.tabBar.items"
        :label="item.title"
        :key="item.fullPath"
        :name="item.fullPath"
        :closable="item.closable"
      ></ElTabPane>
    </ElTabs>
  </div>
</template>

<script lang="ts" setup>
import type { TabsPaneContext, TabPaneName } from "element-plus";
import { useAppStore } from "../stores/AppStore";
import router from "../router";

const appStore = useAppStore();

const onTabChange = (name: TabPaneName) => {
  console.log("onTabChange", name);
};

const onTabClick = (ctx: TabsPaneContext, event: Event) => {
  console.log("onTabClick", ctx, event);
  router.push(ctx.paneName as string);
};

const onTabRemove = (name: TabPaneName) => {
  console.log("onTabRemove", name);
  const index = appStore.tabBar.items.findIndex(
    (tab) => tab.fullPath == (name as string)
  );
  appStore.tabBar.items.splice(index, 1);
};
</script>

<style lang="scss" scoped>
.tab-bar {
  margin: 0.5em 1em;
}
</style>
