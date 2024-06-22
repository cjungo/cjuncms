<template>
  <div class="employee-index-page">
    <div class="employee-info-box">
      <ElForm :model="current" label-width="6em">
        <ElRow>
          <ElCol :span="6">
            <ElFormItem label="ID">
              <ElInput :value="current.id" readonly />
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
      </ElForm>
    </div>
    <div class="employee-list-box">
      <ElAutoResizer>
        <template #default="{ width, height }">
          <VxeTable :width="width" :height="height" :data="rows">
            <VxeColumn type="seq" title="#" width="60" />
            <VxeColumn field="id" title="ID" />
            <VxeColumn field="jobnumber" title="工号" />
            <VxeColumn field="username" title="用户名" />
            <VxeColumn field="fullname" title="全称" />
            <VxeColumn field="nickname" title="昵称" />
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
import { queryEmployee, type Employee } from "../../apis/employee";
import { ElAutoResizer } from "element-plus";
import { Delete, Edit } from "@element-plus/icons-vue";

const isReadonly = ref(true);
const current = ref<Employee>({
  id: 0,
  jobnumber: "",
  username: "",
  nickname: "",
  fullname: "",
  birthday: "",
  avatar_path: "",
  is_removed: 0,
});
const rows = ref<Employee[]>([]);

const onClickEdit = (params: any) => {
  console.log("onClickEdit", params);
};

const onClickDelete = (params: any) => {
  console.log("onClickDelete", params);
};

onBeforeMount(async () => {
  const result = await queryEmployee();
  rows.value = result.data;
});
</script>

<style lang="scss" scoped>
.employee-index-page {
  display: flex;
  flex-direction: column;
  flex-grow: 1;
  padding: 1em;
}

.employee-info-box {
  flex-grow: 0;
}

.employee-list-box {
  flex-grow: 1;
}
</style>
