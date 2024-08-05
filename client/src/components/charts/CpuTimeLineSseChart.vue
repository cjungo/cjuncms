<template>
  <ResizableChart :option="option" @resize="onResize" />
</template>

<script lang="ts" setup>
import { computed, onBeforeMount, ref } from "vue";
import { CjMachineCPUTime, listMachineCpuTimes, watchMachineCpuTimeline } from "../../apis/machine";
import { ResizeParam } from "./ResizableChart.vue";
import { format, parseISO, subDays } from "date-fns";

const cpuTimeItems = ref<CjMachineCPUTime[]>([]);
const resizeParam = ref<ResizeParam>();
const onResize = (param: ResizeParam) => {
  resizeParam.value = param;
};

const option = computed(() => {
  const w = resizeParam.value?.width ?? 0;

  const dateItems = cpuTimeItems.value.map((i) =>
    // format(parseISO(i.create_at), "yyyy-MM-dd HH:mm:ss")
    format(parseISO(i.create_at), "[dd]HH:mm:ss")
  );
  const valueItems = cpuTimeItems.value.map((i) => {
    const used = i.user + i.system;
    const total = used + i.idle;
    return (used / total) * 100;
  });

  return {
    // Make gradient line here
    visualMap: [
      {
        show: false,
        type: "continuous",
        seriesIndex: 0,
        dimension: 0,
        min: 0,
        max: dateItems.length - 1,
      },
    ],
    title: [
      {
        left: "center",
        text: "CPU使用",
      },
    ],
    tooltip: {
      trigger: "axis",
    },
    xAxis: [
      {
        data: dateItems,
      },
    ],
    yAxis: [{}],
    grid: [{}],
    series: [
      {
        type: "line",
        showSymbol: false,
        data: valueItems,
      },
    ],
    backgroundColor: "#fff",
  };
});

onBeforeMount(async () => {
  const cpuTimesResult = await listMachineCpuTimes({
    startAt: subDays(new Date(), 1).toISOString(),
    endAt: new Date().toISOString(),
  });
  cpuTimeItems.value = cpuTimesResult.data;
  watchMachineCpuTimeline(event => {
    const data = JSON.parse(event.data);
    cpuTimeItems.value = data;
  });
});
</script>
