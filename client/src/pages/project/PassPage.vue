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
      <InfoTable
        v-loading="isLoading"
        :loading="isLoading"
        :data="rows"
        @cell-click="onClickCell"
        @scroll="onScroll"
      >
        <VxeColumn fixed="left" type="seq" title="#" width="50" />
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
const isLoading = ref(false);
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

const onScroll: VxeTableEvents.Scroll<CjPass> = async (params) => {
  console.log("onScroll", params);
  const activeTop = params.scrollHeight / 2;
  if (params.scrollTop > activeTop) {
    param.value.skip += rows.value.length;
    await onQuery();
  }
};

const onQuery = async () => {
  if (isLoading.value) return;
  try {
    if (param.value.skip == 0) {
      isLoading.value = true;
    }

    const result = await queryPass(param.value);
    if (param.value.skip == 0) {
      rows.value = result.data;
      current.value = result.data[0];
    } else {
      result.data.forEach((i) => rows.value.push(i));
    }
  } finally {
    isLoading.value = false;
  }
};

onBeforeMount(async () => {
  await onQuery();
});
</script>

<style lang="scss" scoped></style>
