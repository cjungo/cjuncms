<template>
  <CJunCmsPageMainLayout class="project-pass-page">
    <InfoForm>
      <ElRow>
        <ElCol :span="6">
          <ElFormItem label="ID">
            <ElInput :value="current?.id" readonly />
          </ElFormItem>
        </ElCol>
        <ElCol :span="6">
          <ElFormItem label="名称">
            <ElInput :value="current?.title" readonly />
          </ElFormItem>
        </ElCol>
      </ElRow>
    </InfoForm>
    <template #list>
      <InfoTable :data="rows" @cell-click="onClickCell">
        <VxeColumn fixed="left" type="seq" title="#" width="60" />
        <VxeColumn fixed="left" field="id" title="ID" />
        <VxeColumn field="title" title="名称" />
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
      </InfoTable>
    </template>
  </CJunCmsPageMainLayout>
</template>

<script lang="ts" setup>
import { onBeforeMount, ref } from "vue";
import { Delete, Edit } from "@element-plus/icons-vue";
import { queryPass, type CjPass, type QueryPassParam } from "../../apis/pass";
import { VxeTableEvents } from "vxe-table";

const param = ref<QueryPassParam>({
  plain: "",
  skip: 0,
  take: 100,
});
const current = ref<CjPass>();
const rows = ref<CjPass[]>([]);

const onClickCell: VxeTableEvents.CellClick<CjPass> = ({ row, column }) => {
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
  const result = await queryPass(param.value);
  rows.value = result.data;
  current.value = result.data[0];
});
</script>

<style lang="scss" scoped></style>
