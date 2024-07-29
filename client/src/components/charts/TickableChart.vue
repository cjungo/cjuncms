<template>
  <VChart
    ref="vChart"
    class="chart"
    :option="props.option"
    :autoresize="{
      onResize: onResize,
    }"
  />
</template>

<script lang="ts" setup>
import VChart from "vue-echarts";
import { onBeforeMount, onBeforeUnmount, onMounted, ref } from "vue";

export type TickParam = {
  width: number;
  height: number;
};

type VChartType = {
  getWidth(): number;
  getHeight(): number;
};

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

const width = ref(0);
const height = ref(0);

const tick = () => {
  emit("tick", {width: width.value, height: height.value });
};
const tickId = ref(0);

const vChart = ref<VChartType | null>(null);
const onResize = () => {
  console.log("onResize", vChart?.value?.getWidth(), vChart?.value?.getHeight());
  width.value = vChart?.value?.getWidth() ?? 0;
  height.value = vChart?.value?.getHeight() ?? 0;
  tick();
};

onMounted(() => {
  width.value = vChart?.value?.getWidth() ?? 0;
  height.value = vChart?.value?.getHeight() ?? 0;
  console.log("onMounted", vChart?.value?.getWidth(), vChart?.value?.getHeight());
  tick();
});

onBeforeMount(async () => {
  tickId.value = setInterval(tick, props.duration) as any;
});

onBeforeUnmount(() => {
  if (tickId.value != 0) {
    clearInterval(tickId.value);
    tickId.value = 0;
  }
});
</script>
