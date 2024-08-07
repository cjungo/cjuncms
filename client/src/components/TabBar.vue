<template>
  <div class="tab-bar">
    <ElTabs
      v-model="appStore.tabBar.activeName"
      type="card"
      @tab-change="onTabChange"
      @tab-click="onTabClick"
      @tab-remove="onTabRemove"
    >
      <ElTabPane name="/index" :closable="false">
        <template #label>
          <IEpStar />
          <span>首页</span>
        </template>
      </ElTabPane>
      <ElTabPane
        v-for="item in appStore.tabBar.items"
        :key="item.fullPath"
        :name="item.fullPath"
        :closable="item.closable"
      >
        <template #label>
          <!-- <span>{{ metas[item.fullPath].icon }} </span> -->
          <component
            v-if="metas[item.fullPath]?.icon"
            :is="metas[item.fullPath].icon"
          />
          <span>{{ item.title }}</span>
        </template>
      </ElTabPane>
    </ElTabs>
  </div>
</template>

<script lang="ts" setup>
import type { TabsPaneContext, TabPaneName } from "element-plus";
import { useAppStore } from "../stores/AppStore";
import router from "../router";
import { metas } from "../router";

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
  --active-color: #409eff;

  margin: 0.5em 0 0 0;
  padding: 0 1em;
  border-bottom: 1px solid #d8d9df;

  :deep(.el-tabs__header) {
    border: none;
  }

  :deep(.el-tabs__nav-wrap) {
    display: flex !important;
    flex-direction: row;
    justify-content: start;
    align-items: center;
  }

  :deep(.el-tabs__nav-prev), :deep(.el-tabs__nav-next) {
    border: 1px solid #d8d9dd;   
  }
  :deep(.el-tabs__nav-scroll) {
    margin: 0 .5em;
  }

  :deep(.el-tabs__nav) {
    border: none;
  }

  :deep(.el-tabs__item) {
    margin: 0.5em;
    border: 1px solid #d8d9df;
    border-radius: 0.5em;

    & > svg {
      flex-shrink: 0;
      width: 14px;
      height: 14px;
      margin-right: 0.5em;
    }

    &.is-active {
      border: 1px solid var(--active-color);
      border-bottom: 1px solid var(--active-color);

      & i.el-icon {
        border: 1px solid var(--active-color);
      }
    }

    & i.el-icon {
      flex-shrink: 0;
      width: 14px !important;
      height: 14px !important;
      border: 1px solid #d8d9df;
    }
  }
}
</style>
