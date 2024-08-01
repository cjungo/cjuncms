<template>
  <TickableChart :option="option" @tick="tick" />
</template>

<script lang="ts" setup>
import { computed, ref } from "vue";
import { subDays, format, parseISO } from "date-fns";
import {
  type CjMachineCPUTime,
  type ListMachineCpuTimesParam,
  listMachineCpuTimes,
} from "../../apis/machine";
import { type TickParam } from "./TickableChart.vue";

const cpuTimeItems = ref<CjMachineCPUTime[]>([]);

const tickParam = ref<TickParam>();
const option = computed(() => {
  const w = tickParam.value?.width ?? 0;

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

const listParam = ref<ListMachineCpuTimesParam>({
  startAt: subDays(new Date(), 1).toISOString(),
  endAt: new Date().toISOString(),
});
const tick = async (param: TickParam) => {
  const cpuTimesResult = await listMachineCpuTimes(listParam.value);
  cpuTimeItems.value = cpuTimesResult.data;
  console.log("cpu time list.", cpuTimesResult, param);
  tickParam.value = param;
};
</script>
