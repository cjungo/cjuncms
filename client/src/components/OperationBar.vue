<template>
    <div class="operation-bar">
        <slot></slot>
    </div>
</template>

<style lang="scss">
.operation-bar {
  display: flex;
  flex-direction: row;
  padding: 1em;

  background-color: #fff;
  border: 1px solid #d8d9df;
  border-radius: 0.5em;
}
</style>