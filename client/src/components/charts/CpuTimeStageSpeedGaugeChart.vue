<template>
  <TickableChart :option="option" @tick="tick" />
</template>

<script lang="ts" setup>
import { computed, ref } from "vue";
import { CjMachineCPUTime, getMachineCpuTimes } from "../../apis/machine";

const machineCpuTimes = ref<CjMachineCPUTime>({
  id: 0,
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
  guest_nice: 0,
  create_at: "",
});

const option = computed(() => {
  const used = machineCpuTimes.value.user + machineCpuTimes.value.system;
  const total = machineCpuTimes.value.idle + used;
  const percent = (used / total) * 100;
  return {
    series: [
      {
        type: "gauge",
        axisLine: {
          lineStyle: {
            width: 30,
            color: [
              [0.3, "#67e0e3"],
              [0.7, "#37a2da"],
              [1, "#fd666d"],
            ],
          },
        },
        pointer: {
          itemStyle: {
            color: "auto",
          },
        },
        axisTick: {
          distance: -30,
          length: 8,
          lineStyle: {
            color: "#fff",
            width: 2,
          },
        },
        splitLine: {
          distance: -30,
          length: 30,
          lineStyle: {
            color: "#fff",
            width: 4,
          },
        },
        axisLabel: {
          color: "inherit",
          distance: 40,
          fontSize: 20,
        },
        detail: {
          valueAnimation: true,
          formatter: "{value} %",
          color: "inherit",
        },
        data: [
          {
            value: +percent.toFixed(3),
          },
        ],
      },
    ],
  };
});

const tick = async () => {
  const cpuTimesResult = await getMachineCpuTimes();
  machineCpuTimes.value = cpuTimesResult.data;
  console.log("cpu times", cpuTimesResult);
};
</script>
