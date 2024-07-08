<template>
  <VChart class="chart" :option="props.option" :autoresize="true" />
</template>

<script lang="ts" setup>
import VChart from "vue-echarts";
import { onBeforeMount, onBeforeUnmount, ref } from "vue";

const props = withDefaults(
  defineProps<{
    option: any;
    duration?: number;
  }>(),
  {
    duration: 2000,
  }
);
const emit = defineEmits(["tick"]);

const tick = () => emit("tick");
const tickId = ref(0);

onBeforeMount(async () => {
  tick();
  tickId.value = setInterval(tick, props.duration) as any;
});

onBeforeUnmount(() => {
  if (tickId.value != 0) {
    clearInterval(tickId.value);
    tickId.value = 0;
  }
});
</script>
