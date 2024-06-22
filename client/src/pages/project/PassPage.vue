<template>
  <div class="project-pass-page">
    <div class="project-pass-info-box"></div>
    <div class="project-pass-list-box">
      <ElAutoResizer>
        <template #default="{ width, height }">
          <VxeTable :width="width" :height="height" :data="rows"> </VxeTable>
        </template>
      </ElAutoResizer>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { ref } from 'vue';
import { type Pass } from '../../apis/pass';


const rows = ref<Pass[]>([]);

</script>

<style lang="scss" scoped>
.project-pass-page {
  display: flex;
  flex-direction: column;
}

.project-pass-info-box {
  flex-grow: 0;
}

.project-pass-list-box {
  flex-grow: 1;
}
</style>
