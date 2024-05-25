<template>
  <div class="tab-bar">
    <ElTabs v-model="appStore.tabBar.activeName" @tab-click="onTabClick">
      <ElTabPane
        v-for="(item, i) in appStore.tabBar.items"
        :label="item"
        :name="i"
      ></ElTabPane>
    </ElTabs>
  </div>
</template>

<script lang="ts" setup>
import type { TabsPaneContext } from "element-plus";
import { useAppStore } from "../stores/AppStore";

const appStore = useAppStore();

const onTabClick = (tab: TabsPaneContext, event: Event) => {
  console.log(tab, event);
};
</script>
