<template>
    <div class="top-bar">
        <div class="top-title">

        </div>
        <ElMenu mode="horizontal" :default-active="appStore.topBar.defaultActive">
            <ElMenuItem></ElMenuItem>
        </ElMenu>
    </div>
</template>

<script lang="ts" setup>
import { reactive } from "vue";
import { useAppStore } from "../stores/AppStore";

export type TopBarItem = {
    title: string;
    permissions?: [string];
};

const appStore = useAppStore();

const items = reactive<[TopBarItem]>([
    {
        title: "首页"
    }
]);

</script>

<style lang="scss" scoped>
.top-bar {
    display: flex;
    flex-direction: column;
}

.top-title {
    width: 200px;
    height: 100%;
}
</style>
