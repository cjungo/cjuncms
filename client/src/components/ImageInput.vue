<template>
  <div class="image-input">
    <input @input="onInput" type="file" style="opacity: 0" />
    <img v-if="displayUrlStore" :src="displayUrlStore" alt="image" />
    <div class="close-button" @click="onClose">
      <IEpCloseBold />
    </div>
  </div>
</template>

<script lang="ts" setup>
import { computed, ref } from "vue";

// const props = withDefaults(defineProps<{
//     modelValue:
// }>(), {

// });

const emits = defineEmits(["input"]);

const displayUrlStore = ref<string>();
const currentFileStore = ref<File>();

const currentFile = computed({
  get: () => {
    return currentFileStore.value;
  },
  set: (v) => {
    if (displayUrlStore.value && displayUrlStore.value.startsWith("blob:")) {
      URL.revokeObjectURL(displayUrlStore.value);
    }

    displayUrlStore.value = v ? URL.createObjectURL(v) : undefined;
    currentFileStore.value = v;
    emits("input", v);
  },
});

const onInput = (event: Event) => {
  const element = event.target as HTMLInputElement;
  currentFile.value = element.files![0];
  element.value = "";
};

const onClose = () => {
  currentFile.value = undefined;
};
</script>

<style lang="scss" scoped>
.image-input {
  position: relative;
  display: flex;
  width: 100%;
  aspect-ratio: 1;
  border: 1px solid #ddd;
  border-radius: 1em;
  margin: 0;
  padding: 0;
  overflow: hidden;

  & > input {
    position: absolute;
    z-index: 100;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    width: 100%;
    height: 100%;
  }

  & > image {
    position: absolute;
    z-index: 10;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    width: 100%;
    height: 100%;
  }

  & > .close-button {
    display: flex;
    justify-content: center;
    align-items: center;
    position: absolute;
    z-index: 1000;
    top: 0.5em;
    right: 0.5em;
    width: 20%;
    height: 20%;
    border-radius: 50%;
    background-color: #4444;
    color: #fff;
  }
}
</style>
