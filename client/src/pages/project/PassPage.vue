<template>
  <div class="project-pass-page">
    <div class="project-pass-info-box"></div>
    <div class="project-pass-list-box">
      <ElAutoResizer>
        <template #default="{ width, height }">
          <VxeTable :width="width" :height="height" :data="rows">
            <VxeColumn type="seq" title="#" width="60" />
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
import { queryPass, type CjPass, type QueryPassParam } from "../../apis/pass";

const param = ref<QueryPassParam>({
  plain: "",
  skip: 0,
  take: 100,
});
const rows = ref<CjPass[]>([]);

const onClickEdit = (params: any) => {
  console.log("onClickEdit", params);
};

const onClickDelete = (params: any) => {
  console.log("onClickDelete", params);
};

onBeforeMount(async () => {
  const result = await queryPass(param.value);
  rows.value = result.data;
});
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
