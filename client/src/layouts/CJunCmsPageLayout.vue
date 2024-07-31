<template>
  <div class="cjuncms-page-layout">
    <LeftBar />
    <div class="cjuncms-page-content">
      <TopBar />
      <div class="cjuncms-page-main">
        <TabBar/>
        <RouterView v-slot="{ Component }">
          <KeepAlive>
            <component :is="Component" />
          </KeepAlive>
        </RouterView>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts"></script>

<style lang="scss" scoped>
.cjuncms-page-layout {
  display: flex;
  flex-direction: row;

  width: 100%;
  height: 100%;
  overflow: hidden;

  // background-color: red;
}

.cjuncms-page-content {
  display: flex;
  flex-direction: column;
  overflow: hidden;
  flex-grow: 1;

  // background-color: blue;
}

.cjuncms-page-main {
  display: flex;
  overflow: hidden;
  flex-direction: column;
  flex-grow: 1;
}
</style>
