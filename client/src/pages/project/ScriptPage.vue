<template>
  <div class="project-script-page">
    <ElAutoResizer>
      <template #default="{ width, height }">
        <VxeTable :width="width" :height="height" :data="rows">
          <VxeColumn type="seq" title="#" width="60" />
          <VxeColumn field="id" title="ID" />
          <VxeColumn field="title" title="标题" />
          <VxeColumn field="content" title="内容" />
          <VxeColumn>
            <ElButtonGroup>
              <ElButton
                @click="onClickEdit"
                type="primary"
                :icon="Edit"
              ></ElButton>
              <ElButton
                @click="onClickDelete"
                type="danger"
                :icon="Delete"
              ></ElButton>
            </ElButtonGroup>
          </VxeColumn>
        </VxeTable>
      </template>
    </ElAutoResizer>
  </div>
</template>

<script lang="ts" setup>
import { onBeforeMount, ref } from "vue";
import { CjScript, queryScript } from "../../apis/script";
import { Delete, Edit } from "@element-plus/icons-vue";

const rows = ref<CjScript[]>([]);

const onClickEdit = (params: any) => {
  console.log("onClickEdit", params);
};

const onClickDelete = (params: any) => {
  console.log("onClickDelete", params);
};

onBeforeMount(async () => {
  const response = await queryScript();
  rows.value = response.data;
});
</script>

<style lang="scss" scoped>
.project-script-page {
  display: flex;
  flex-direction: column;
}
</style>
