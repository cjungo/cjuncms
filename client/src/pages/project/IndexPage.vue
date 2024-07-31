<template>
  <CJunCmsPageMainLayout>
    <ElForm>
      <ElRow>
        <ElCol :span="6">
          <ElFormItem label="ID">
            <ElInput :value="current?.id" readonly />
          </ElFormItem>
        </ElCol>
        <ElCol :span="6">
          <ElFormItem label="项目名">
            <ElInput :value="current?.name" readonly />
          </ElFormItem>
        </ElCol>
      </ElRow>
    </ElForm>
    <template #list>
      <ElAutoResizer>
        <template #default="{ width, height }">
          <VxeTable
            border
            show-overflow
            :column-config="{ resizable: true }"
            :width="width"
            :height="height"
            :data="rows"
          >
            <VxeColumn fixed="left" type="seq" title="#" width="60" />
            <VxeColumn fixed="left" field="id" title="ID" />
            <VxeColumn field="name" title="项目名" />
            <VxeColumn fixed="right">
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
    </template>
  </CJunCmsPageMainLayout>
</template>

<script lang="ts" setup>
import { onBeforeMount, ref } from "vue";
import { Delete, Edit } from "@element-plus/icons-vue";
import { queryProject, type CjProject } from "../../apis/project";

const current = ref<CjProject>();
const rows = ref<CjProject[]>([]);

const onClickEdit = (params: any) => {
  console.log("onClickEdit", params);
};

const onClickDelete = (params: any) => {
  console.log("onClickDelete", params);
};

onBeforeMount(async () => {
  const result = await queryProject();
  rows.value = result.data;
});
</script>

<style lang="scss" scoped></style>
