import { defineStore } from "pinia";
import { reactive } from "vue";

export type DialogState = {
  visible: boolean;
};

export type TabBarItem = {
  title: string;
  url: string;
};

export type TabBarState = {
  activeName: string;
  items: Array<TabBarItem>;
};

export type LeftBarState = {
  defaultActive: string;
  isCollapse: boolean;
};

export type TopBarState = {
  defaultActive: string;
};

export const useAppStore = defineStore(
  "app",
  () => {
    const dialog = reactive<{ [k: string]: DialogState }>({
      logout: {
        visible: false,
      },
    });
    const tabBar = reactive<TabBarState>({
      activeName: "home",
      items: [],
    });
    const leftBar = reactive<LeftBarState>({
      defaultActive: "setting",
      isCollapse: false,
    });
    const topBar = reactive<TopBarState>({
      defaultActive: "setting",
    });

    return {
      dialog,
      tabBar,
      leftBar,
      topBar,
    };
  },
  {
    persist: true,
  }
);
