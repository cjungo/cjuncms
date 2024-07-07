<template>
  <div class="index-page">
    <v-chart class="chart" :option="option" :autoresize="true" />
  </div>
</template>

<script setup lang="ts">
import { use } from "echarts/core";
import { CanvasRenderer } from "echarts/renderers";
import { PieChart } from "echarts/charts";
import {
  TitleComponent,
  TooltipComponent,
  LegendComponent,
} from "echarts/components";
import VChart, { THEME_KEY } from "vue-echarts";
import { ref, provide, onBeforeMount, reactive, computed } from "vue";
import {
  getMachineCpuInfo,
  type MachineCpuInfo,
  getMachineCpuTimes,
  type MachineCpuTime,
  getMachineVirtualMemory,
  type MachineVirtualMemory,
} from "../apis/machine";
import { onBeforeRouteLeave, onBeforeRouteUpdate } from "vue-router";

use([
  CanvasRenderer,
  PieChart,
  TitleComponent,
  TooltipComponent,
  LegendComponent,
]);

const machineCpuInfo = ref<MachineCpuInfo>({
  cpu: 0,
  vendorId: "",
  family: "",
  model: "",
  stepping: 0,
  physicalId: "",
  coreId: "",
  cores: 0,
  modelName: "",
  mhz: 0,
  cacheSize: 0,
  flags: [],
  microcode: "",
});
const machineCpuTimes = ref<MachineCpuTime>({
  cpu: "",
  user: 0,
  system: 0,
  idle: 0,
  nice: 0,
  iowait: 0,
  irq: 0,
  softirq: 0,
  steal: 0,
  guest: 0,
  guestNice: 0,
});
const machineVirtualMemory = ref<MachineVirtualMemory>({
  total: 0,
  available: 0,
  used: 0,
  usedPercent: 0,
  free: 0,
  active: 0,
  inactive: 0,
  wired: 0,
  laundry: 0,
  buffers: 0,
  cached: 0,
  writeBack: 0,
  dirty: 0,
  writeBackTmp: 0,
  shared: 0,
  slab: 0,
  sreclaimable: 0,
  sunreclaim: 0,
  pageTables: 0,
  swapCached: 0,
  commitLimit: 0,
  committedAS: 0,
  highTotal: 0,
  highFree: 0,
  lowTotal: 0,
  lowFree: 0,
  swapTotal: 0,
  swapFree: 0,
  mapped: 0,
  vmallocTotal: 0,
  vmallocUsed: 0,
  vmallocChunk: 0,
  hugePagesTotal: 0,
  hugePagesFree: 0,
  hugePagesRsvd: 0,
  hugePagesSurp: 0,
  hugePageSize: 0,
  anonHugePages: 0,
});

// provide(THEME_KEY, "dark");

const option = computed(() => {
  const cpu = machineCpuTimes.value ?? {
    system: 0,
    user: 0,
    idle: 0,
  };

  return {
    title: {
      text: "服务状态",
      left: "center",
    },
    tooltip: {
      trigger: "item",
      formatter: "{a} <br/>{b} : {c} ({d}%)",
    },
    legend: {
      orient: "vertical",
      left: "4%",
      data: [
        "CPU 系统占用",
        "CPU 用户占用",
        "CPU 空闲",
        "内存占用",
        "内存空闲",
      ],
    },
    series: [
      {
        name: "CPU",
        type: "pie",
        radius: ["10%", "24%"],
        center: ["30%", "40%"],
        data: [
          { value: cpu.system, name: "CPU 系统占用" },
          { value: cpu.user, name: "CPU 用户占用" },
          { value: cpu.idle, name: "CPU 空闲" },
        ],
        emphasis: {
          itemStyle: {
            shadowBlur: 10,
            shadowOffsetX: 0,
            shadowColor: "rgba(0, 0, 0, 0.5)",
          },
        },
      },
      {
        name: "内存",
        type: "pie",
        radius: ["10%", "24%"],
        center: ["70%", "40%"],
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
  const cpuTimesResult = await getMachineCpuTimes();
  machineCpuTimes.value = cpuTimesResult.data;
  console.log("cpu times", cpuTimesResult);

  const virtualMemoryResult = await getMachineVirtualMemory();
  machineVirtualMemory.value = virtualMemoryResult.data;
  console.log("virtual memory", virtualMemoryResult);
};
const tickId = ref(0);

onBeforeMount(async () => {
  const cpuInfoResult = await getMachineCpuInfo();
  machineCpuInfo.value = cpuInfoResult.data;
  console.log("cpu", cpuInfoResult);

  tickId.value = setInterval(tick, 2000) as any;
});

onBeforeRouteUpdate(async () => {
  if (tickId.value == 0) {
    tickId.value = setInterval(tick, 2000) as any;
  }
});

onBeforeRouteLeave(async () => {
  if (tickId.value != 0) {
    clearInterval(tickId.value);
    tickId.value = 0;
  }
});
</script>

<script lang="ts">
export default {
  name: "",
  beforeRouteEnter() {},
};
</script>

<style lang="scss" scoped>
.index-page {
  width: 100%;
  flex-grow: 1;
  overflow: hidden;
}
</style>
