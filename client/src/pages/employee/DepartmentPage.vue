<template>
  <CJunCmsPageMainLayout class="employee-department-page">
    <OperationBar> </OperationBar>
    <template #list>
      <ElAutoResizer>
        <template #default="{ width, height }">
          <ElTreeV2
            class="tree"
            show-checkbox
            :width="width"
            :height="height"
            :data="tree"
            :props="{ value: 'id', label: 'title', children: 'children' }"
            :item-size="50"
            :expand-on-click-node="false"
            @check-change="onCheckChange"
          >
            <template #default="{ node }">
              <div class="node-row">
                <ElInput :value="node.label" />
                <div class="node-operation-bar">
                  <ElButton
                    type="success"
                    :icon="Plus"
                    circle
                    @click="onClickAdd(node)"
                  />
                  <ElButton
                    type="primary"
                    :icon="EditPen"
                    circle
                    @click="onClickEdit(node)"
                  />
                  <ElButton
                    type="danger"
                    :icon="Delete"
                    circle
                    @click="onClickDrop(node)"
                  />
                </div>
              </div>
            </template>
          </ElTreeV2>
        </template>
      </ElAutoResizer>
    </template>
  </CJunCmsPageMainLayout>
</template>

<script lang="ts" setup>
import { onBeforeMount, ref } from "vue";
import {
  type DepartmentNode,
  type QueryDepartmentParam,
  queryDepartment,
} from "../../apis/department";
import { ElAutoResizer } from "element-plus";
import { type TreeNodeData } from "element-plus/lib/components/tree-v2/src/types.js";
import { type TreeKey } from "element-plus/es/components/tree-v2/src/types.mjs";
import { Plus, EditPen, Delete } from "@element-plus/icons-vue";

type Node = {
  children?: Array<Node>;
  data: DepartmentNode;
  key: TreeKey;
  isLeaf: boolean;
  label: string;
  level: number;
  parent?: Node;
};

const param = ref<QueryDepartmentParam>({});
const tree = ref<Array<DepartmentNode>>([]);

const onCheckChange = (data: TreeNodeData, checked: boolean) => {
  console.log("onCheckChange", data, checked);
};

const onClickAdd = (node: Node) => {
  console.log("onClickAdd", node);
};
const onClickEdit = (node: Node) => {
  console.log("onClickEdit", node);
}
const onClickDrop = (node: Node) => {
  console.log("onClickDrop", node);
};

onBeforeMount(async () => {
  const response = await queryDepartment(param.value);
  tree.value = response.data;
});
</script>

<style lang="scss" scoped>
.tree {
  :deep(.el-vl__wrapper) {
    padding: 1em;
  }
}
.node-row {
  display: flex;
  flex-grow: 1;
  flex-direction: row;
  align-items: center;
  justify-content: space-between;
  overflow: hidden;
  padding: 0.5em 0.5em .5em 1em;
  border: 1px solid var(--el-border-color);
  border-radius: var(--el-border-radius-round);
}

.node-operation-bar {
  padding-left: 1em;
}
</style>
