<template>
  <div class="tab-bar">
    <ElTabs
      v-model="appStore.tabBar.activeName"
      @tab-change="onTabChange"
      @tab-click="onTabClick"
      @tab-remove="onTabRemove"
    >
      <ElTabPane label="首页"></ElTabPane>
      <ElTabPane
        v-for="(item, i) in appStore.tabBar.items"
        :label="item.title"
        :index="item.fullPath"
        :name="i"
        :closable="item.closable"
      ></ElTabPane>
    </ElTabs>
  </div>
</template>

<script lang="ts" setup>
import type { TabsPaneContext, TabPaneName } from "element-plus";
import { useAppStore } from "../stores/AppStore";

const appStore = useAppStore();

const onTabChange = (name: TabPaneName) => {
  console.log("onTabChange", name);
};

const onTabClick = (tab: TabsPaneContext, event: Event) => {
  console.log("onTabClick", tab, event);
};

const onTabRemove = (name: TabPaneName) => {
  console.log("onTabRemove", name);
};
</script>

<style lang="scss" scoped>
.tab-bar {
  margin: .5em 1em;
}
</style>