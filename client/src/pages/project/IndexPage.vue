<template>
  <div class="project-index-page">
    <div class="project-info-box"></div>
    <div class="project-list-box">
      <ElAutoResizer>
        <template #default="{ width, height }">
          <VxeTable :width="width" :height="height" :data="rows">
            <VxeColumn type="seq" title="#" width="60" />
            <VxeColumn field="id" title="ID" />
            <VxeColumn field="name" title="项目名" />
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
  </div>
</template>

<script lang="ts" setup>
import { onBeforeMount, ref } from "vue";
import { Delete, Edit } from "@element-plus/icons-vue";
import { queryProject, type Project } from "../../apis/project";

const rows = ref<Project[]>([]);

const onClickEdit = (params: any) => {
  console.log("onClickEdit", params);
};

const onClickDelete = (params: any) => {
  console.log("onClickDelete", params);
};

onBeforeMount(async() => {
    const result = await queryProject();
    rows.value = result.data;
});
</script>

<style lang="scss" scoped>
.project-index-page {
  display: flex;
  flex-direction: column;
}
</style>
