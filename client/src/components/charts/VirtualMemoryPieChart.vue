<template>
  <TickableChart :option="option" @tick="tick" />
</template>

<script lang="ts" setup>
import { ref, computed } from "vue";
import {
  type CjMachineVirtualMemory,
  getMachineVirtualMemory,
} from "../../apis/machine";

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
  return {
    title: {
      text: "内存状态",
      left: "center",
    },
    tooltip: {
      trigger: "item",
      formatter: "{a} <br/>{b} : {c} ({d}%)",
    },
    legend: {
      orient: "vertical",
      left: "4%",
      data: ["内存占用", "内存空闲"],
    },
    series: [
      {
        name: "内存",
        type: "pie",
        radius: ["10%", "40%"],
        center: ["50%", "50%"],
        data: [
          { value: machineVirtualMemory.value.used, name: "内存占用" },
          { value: machineVirtualMemory.value.free, name: "内存空闲" },
        ],
        emphasis: {
          itemStyle: {
            shadowBlur: 10,
            shadowOffsetX: 0,
            shadowColor: "rgba(0, 0, 0, 0.5)",
          },
        },
      },
    ],
  };
});

const tick = async () => {
  const virtualMemoryResult = await getMachineVirtualMemory();
  machineVirtualMemory.value = virtualMemoryResult.data;
  console.log("virtual memory", virtualMemoryResult);
};
</script>
