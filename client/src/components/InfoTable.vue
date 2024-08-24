<template>
  <ElAutoResizer>
    <template #default="{ width, height }">
      <VxeTable
        ref="tableRef"
        border
        round
        show-overflow
        :column-config="{ resizable: true }"
        :checkbox-config="{ highlight: true, range: true }"
        :width="width"
        :height="height"
        v-bind="$attrs"
        @checkbox-change="selectChangeEvent"
      >
        <VxeColumn
          fixed="left"
          type="checkbox"
          width="50"
          align="center"
          header-align="center"
        >
          <template #header="{ checked, indeterminate }">
            <VxeCheckbox
              :value="checked"
              :indeterminate="indeterminate"
              @native.click.stop="toggleAllCheckboxEvent"
            />
          </template>
          <template #checkbox="{ row, checked }">
            <VxeCheckbox
              :value="checked"
              @click.native.stop="toggleCheckboxEvent(row)"
            />
          </template>
        </VxeColumn>
        <slot></slot>
      </VxeTable>
    </template>
  </ElAutoResizer>
</template>

<script lang="ts" setup>
import { ref } from "vue";
import { VxeTableEvents, type VxeTableInstance } from "vxe-table";

const emit = defineEmits(["select"]);

const tableRef = ref<VxeTableInstance>();

const toggleAllCheckboxEvent = () => {
  const $table = tableRef.value;
  if ($table) {
    $table.toggleAllCheckboxRow();
    console.log("toggleAllCheckboxEvent");
  }
};

const toggleCheckboxEvent = (row: any) => {
  const $table = tableRef.value;
  if ($table) {
    $table.toggleCheckboxRow(row);
    console.log("toggleCheckboxEvent");
  }
};

const selectChangeEvent: VxeTableEvents.CheckboxChange<any> = ({ $table }) => {
  const records = $table.getCheckboxRecords();
  emit("select", records);
};
</script>

<style lang="scss" scoped>
// :deep(.col--checkbox .vxe-cell .vxe-cell--title) {
//   display: flex;
//   justify-content: center;
//   align-items: center;
//   padding: 0;
// }
// :deep(.el-checkbox) {
//   display: flex;
//   justify-content: center;
//   align-items: center;
//   padding: 0;
// }
</style>
