import { defineStore } from "pinia";
import { ref } from "vue";

export type DialogState = {
  visible: boolean;
};

export type TabBarItem = {
  title: string;
  fullPath: string;
  closable: boolean;
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
    const dialog = ref<{ [k: string]: DialogState }>({
      logout: {
        visible: false,
      },
    });
    const tabBar = ref<TabBarState>({
      activeName: "home",
      items: [],
    });
    const leftBar = ref<LeftBarState>({
      defaultActive: "setting",
      isCollapse: false,
    });
    const topBar = ref<TopBarState>({
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
