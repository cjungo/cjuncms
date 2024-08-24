<template>
  <CJunCmsPageMainLayout class="project-index-page">
    <InfoForm :model="current">
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
    </InfoForm>
    <template #list>
      <InfoTable :data="rows" @cell-click="onClickCell">
        <VxeColumn fixed="left" field="id" title="ID" width="60" />
        <VxeColumn field="name" title="项目名" width="120"/>
        <VxeColumn fixed="right" width="120">
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
      </InfoTable>
    </template>
  </CJunCmsPageMainLayout>
</template>

<script lang="ts" setup>
import { onBeforeMount, ref } from "vue";
import { Delete, Edit } from "@element-plus/icons-vue";
import { queryProject, QueryProjectParam, type CjProject } from "../../apis/project";
import { VxeTableEvents } from "vxe-table";

const current = ref<CjProject>();
const rows = ref<CjProject[]>([]);
const param = ref<QueryProjectParam>({
  skip: 0,
  take: 0,
  plain: ""
});

const onClickCell: VxeTableEvents.CellClick<CjProject> = ({ row, column }) => {
  current.value = row;
  console.log("onClickCell", row, column);
};

const onClickEdit = (params: any) => {
  console.log("onClickEdit", params);
};

const onClickDelete = (params: any) => {
  console.log("onClickDelete", params);
};

onBeforeMount(async () => {
  const result = await queryProject(param.value);
  rows.value = result.data;
  current.value = result.data[0];
});
</script>

<style lang="scss" scoped>
.project-index-page {
}
</style>
