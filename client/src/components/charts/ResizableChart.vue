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
import { onMounted, ref } from "vue";
import VChart from "vue-echarts";

type VChartType = {
  getWidth(): number;
  getHeight(): number;
};

export type ResizeParam = {
  width: number;
  height: number;
};

const props = withDefaults(
  defineProps<{
    option: any;
  }>(),
  {}
);
const emit = defineEmits(["resize"]);

const width = ref(0);
const height = ref(0);

const vChart = ref<VChartType | null>(null);
const onResize = () => {
  console.log(
    "onResize",
    vChart?.value?.getWidth(),
    vChart?.value?.getHeight()
  );
  width.value = vChart?.value?.getWidth() ?? 0;
  height.value = vChart?.value?.getHeight() ?? 0;
  emit("resize", { width: width.value, height: height.value });
};

onMounted(() => {
    onResize();
});
</script>
