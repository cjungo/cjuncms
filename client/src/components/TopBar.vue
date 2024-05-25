<template>
    <div class="top-bar">
        <div class="top-title">

        </div>
        <ElMenu mode="horizontal" :default-active="appStore.topBar.defaultActive">

        </ElMenu>
    </div>
</template>

<script lang="ts" setup>
import { useAppStore } from "../stores/AppStore";

export type TopBarItem = {
    title: string;
    permissions?: [string];
};

const appStore = useAppStore();

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
