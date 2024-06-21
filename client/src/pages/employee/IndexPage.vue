<template>
  <div class="employee-index-page">
    <div class="employee-info-box">
      <ElForm :model="current" label-width="auto">
        <ElRow>
          <ElCol :span="4">
            <ElFormItem label="ID">
              <ElInput :value="current.id" readonly />
            </ElFormItem>
          </ElCol>
          <ElCol :span="4">
            <ElFormItem label="用户名">
              <ElInput v-model="current.username" :readonly="isReadonly" />
            </ElFormItem>
          </ElCol>
          <ElCol :span="4">
            <ElFormItem lable="工号">
              <ElInput v-model="current.jobnumber" :readonly="isReadonly" />
            </ElFormItem>
          </ElCol>
          <ElCol :span="4">
            <ElFormItem lable="昵称">
              <ElInput v-model="current.nickname" :readonly="isReadonly" />
            </ElFormItem>
          </ElCol>
        </ElRow>
      </ElForm>
    </div>
    <div class="employee-list-box">
      <ElAutoResizer>
        <template #default="{ width, height }">
          <ElTableV2
            :columns="columns"
            :data="rows"
            :width="width"
            :height="height"
            :border="true"
          >
          </ElTableV2>
        </template>
      </ElAutoResizer>
    </div>
  </div>
</template>

<script lang="tsx" setup>
import { onBeforeMount, ref } from "vue";
import { queryEmployee, type Employee } from "../../apis/employee";
import { TableV2FixedDir, ElButton, type Column } from "element-plus";

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

const columns = ref<Column[]>([
  // {
  //   key: "id",
  //   dataKey: "id",
  //   title: "id",
  //   width: 100,
  //   fixed: TableV2FixedDir.LEFT,
  // },
  {
    key: "jobnumber",
    dataKey: "jobnumber",
    title: "jobnumber",
    width: 100,
  },
  {
    key: "username",
    dataKey: "username",
    title: "username",
    width: 100,
  },
  {
    key: "nickname",
    dataKey: "nickname",
    title: "nickname",
    width: 100,
  },
  {
    title: "操作",
    width: 200,
    fixed: TableV2FixedDir.RIGHT,
    cellRenderer: (params) => {
      return (
        <>
          <ElButton type="danger" onClick={() => onClickDelete(params)}>
            <span>删除</span>
          </ElButton>
        </>
      );
    },
  },
]);

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
