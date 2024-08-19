<template>
    <ElAutoResizer>
        <template #default="{ width, height }">
          <VxeTable
            border
            show-overflow
            :column-config="{ resizable: true }"
            :width="width"
            :height="height"
            v-bind="$attrs"
          >
            <slot></slot>
          </VxeTable>
        </template>
      </ElAutoResizer>
</template>