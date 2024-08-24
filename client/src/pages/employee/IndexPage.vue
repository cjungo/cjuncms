<template>
  <CJunCmsPageMainLayout class="employee-index-page">
    <InfoForm :model="current">
      <ElRow>
        <ElCol :span="6">
          <ElFormItem label="ID">
            <ElInput :value="current?.id" readonly />
          </ElFormItem>
        </ElCol>
        <ElCol :span="6">
          <ElFormItem label="用户名">
            <ElInput v-model="current.username" :readonly="isReadonly" />
          </ElFormItem>
        </ElCol>
        <ElCol :span="6">
          <ElFormItem label="工号">
            <ElInput v-model="current.jobnumber" :readonly="isReadonly" />
          </ElFormItem>
        </ElCol>
        <ElCol :span="6">
          <ElFormItem label="昵称">
            <ElInput v-model="current.nickname" :readonly="isReadonly" />
          </ElFormItem>
        </ElCol>
      </ElRow>
    </InfoForm>
    <template #list>
      <InfoTable v-loading="isLoading" :data="rows" @cell-click="onClickCell">
        <VxeColumn fixed="left" field="id" title="ID" />
        <VxeColumn field="jobnumber" title="工号" />
        <VxeColumn field="username" title="用户名" />
        <VxeColumn field="fullname" title="全称" />
        <VxeColumn field="nickname" title="昵称" />
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
import {
  queryEmployee,
  type QueryEmployeeParam,
  type CjEmployee,
} from "../../apis/employee";
import { Delete, Edit } from "@element-plus/icons-vue";
import { VxeTableEvents } from "vxe-table";

const isReadonly = ref(true);
const isLoading = ref(false);
const isAll = ref(false);
const current = ref<CjEmployee>({
  id: 0,
  jobnumber: "",
  username: "",
  nickname: "",
  is_removed: 0,
});
const rows = ref<CjEmployee[]>([]);
const param = ref<QueryEmployeeParam>({
  skip: 0,
  take: 0,
  plain: "",
});

const onClickCell: VxeTableEvents.CellClick<CjEmployee> = ({ row, column }) => {
  current.value = row;
  console.log("onClickCell", row, column);
};

const onClickEdit = (params: any) => {
  console.log("onClickEdit", params);
};

const onClickDelete = (params: any) => {
  console.log("onClickDelete", params);
};

const onQuery = async () => {
  if (isLoading.value) return;
  try {
    isLoading.value = true;
    const result = await queryEmployee(param.value);
    rows.value = result.data;
    if (param.value?.skip == 0) {
      current.value = result.data[0];
    }
  } finally {
    isLoading.value = false;
  }
};

onBeforeMount(async () => {
  await onQuery();
});
</script>

<style lang="scss" scoped>
.employee-index-page {
}
</style>
