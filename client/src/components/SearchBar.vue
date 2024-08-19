<template>
  <div class="search-bar">
    <IEpSearch class="search-bar-icon" @click="onClickSearch"/>
    <input v-model="plain" />
  </div>
</template>

<script lang="ts" setup>
import { ref } from "vue";

const plain = ref("");

const onClickSearch = () => {

}
</script>

<style lang="scss">
.search-bar {
  display: flex;
  padding: .5em;
  border: 1px solid #efeef4;
  border-radius: .8em;

  & > .search-bar-icon {
    width: 50%;
    height: 50%;
    color: #444;
    cursor: pointer;
  }

  & > input {
    border: none;
    outline: none;
  }
}
</style>
