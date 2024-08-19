<template>
  <ElForm v-bind="$attrs" label-width="6em" class="info-form">
    <slot></slot>
  </ElForm>
</template>

<style lang="scss" scoped>
.info-form {
  padding: 1em;

  background-color: #fff;
  border: 1px solid #d8d9df;
  border-radius: .5em;
}
</style>