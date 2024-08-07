<template>
  <ResizableChart :option="option" @resize="onResize" />
</template>

<script lang="ts" setup>
import { computed, onBeforeMount, ref } from "vue";
import { ResizeParam } from "./ResizableChart.vue";
import {
  CjMachineVirtualMemory,
  getMachineVirtualMemory,
  watchMachineVirtualMemory,
} from "../../apis/machine";

const resizeParam = ref<ResizeParam>();
const onResize = (param: ResizeParam) => {
  resizeParam.value = param;
};
const machineVirtualMemory = ref<CjMachineVirtualMemory>({
  id: 0,
  total: 0,
  available: 0,
  used: 0,
  used_percent: 0,
  free: 0,
  active: 0,
  inactive: 0,
  wired: 0,
  laundry: 0,
  buffers: 0,
  cached: 0,
  write_back: 0,
  dirty: 0,
  write_back_tmp: 0,
  shared: 0,
  slab: 0,
  sreclaimable: 0,
  sunreclaim: 0,
  page_tables: 0,
  swap_cached: 0,
  commit_limit: 0,
  committed_as: 0,
  high_total: 0,
  high_free: 0,
  low_total: 0,
  low_free: 0,
  swap_total: 0,
  swap_free: 0,
  mapped: 0,
  vmalloc_total: 0,
  vmalloc_used: 0,
  vmalloc_chunk: 0,
  huge_pages_total: 0,
  huge_pages_free: 0,
  huge_pages_rsvd: 0,
  huge_pages_surp: 0,
  huge_page_size: 0,
  anon_huge_pages: 0,
  create_at: "",
});
const option = computed(() => {
  const w = resizeParam.value?.width ?? 0;
  const percent = (
    Number(
      (
        machineVirtualMemory.value.used / machineVirtualMemory.value.total
      ).toFixed(4)
    ) * 100
  ).toFixed(2);
  // console.log('percent', percent, machineVirtualMemory.value.used, machineVirtualMemory.value.total);
  return {
    title: {
      text: "内存状态",
      left: "center",
      textStyle: {
        fontSize: w * 0.08,
      },
    },
    series: [
      {
        type: "gauge",
        progress: {
          show: true,
          width: w * 0.05,
        },
        axisLine: {
          lineStyle: {
            width: w * 0.05,
          },
        },
        axisTick: {
          show: false,
        },
        splitLine: {
          length: w * 0.02,
          distance: w * 0.007,
          lineStyle: {
            width: w * 0.005,
            color: "#999",
          },
        },
        axisLabel: {
          distance: w * 0.06,
          color: "#999",
          fontSize: w * 0.05,
        },
        anchor: {
          show: true,
          showAbove: true,
          size: w * 0.05,
          itemStyle: {
            borderWidth: w * 0.01,
          },
        },
        title: {
          show: false,
        },
        detail: {
          valueAnimation: true,
          fontSize: w * 0.125,
          offsetCenter: [0, w * 0.3],
        },
        data: [
          {
            value: percent,
          },
        ],
      },
    ],
  };
});

onBeforeMount(async () => {
  const virtualMemoryResult = await getMachineVirtualMemory();
  machineVirtualMemory.value = virtualMemoryResult.data;
  watchMachineVirtualMemory((message) => {
    const data = JSON.parse(message.data);
    machineVirtualMemory.value = data;
  });
});
</script>
