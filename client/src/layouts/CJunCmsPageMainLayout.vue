<template>
  <div class="cjuncms-page-main-layout">
    <div class="cjuncms-page-main-info-box">
        <slot></slot>
    </div>
    <div class="cjuncms-page-main-list-box">
        <slot name="list"></slot>
    </div>
  </div>
</template>

<script lang="ts" setup>
</script>

<style lang="scss" scoped>
.cjuncms-page-main-layout {
  display: flex;
  overflow: hidden;
  flex-grow: 1;
  flex-direction: column;
  background-color: #f6f8f9;
}

.cjuncms-page-main-info-box {
    overflow: hidden;
    flex-shrink: 0;
    padding: 1ex;
    max-height: 100%;
}

.cjuncms-page-main-list-box {
  flex-grow: 1;
  overflow: hidden;
  padding: 1ex;
}
</style>
