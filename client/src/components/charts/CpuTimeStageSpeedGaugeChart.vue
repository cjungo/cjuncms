<template>
  <TickableChart :option="option" @tick="tick" />
</template>

<script lang="ts" setup>
import { computed, ref } from "vue";
import { CjMachineCPUTime, getMachineCpuTimes } from "../../apis/machine";
import { type TickParam } from "./TickableChart.vue";

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

const tickParam = ref<TickParam>();
const option = computed(() => {
  const w = tickParam.value?.width ?? 0;

  const used = machineCpuTimes.value.user + machineCpuTimes.value.system;
  const total = machineCpuTimes.value.idle + used;
  const percent = (used / total) * 100;
  return {
    title: {
      text: "CPU",
      left: "center",
    },
    series: [
      {
        type: "gauge",
        axisLine: {
          lineStyle: {
            width: w * 0.06,
            color: [
              [0.3, "#67e0e3"],
              [0.7, "#37a2da"],
              [1, "#fd666d"],
            ],
          },
        },
        pointer: {
          // length: 100,
          itemStyle: {
            color: "auto",
          },
        },
        axisTick: {
          distance: -(w * 0.06),
          length: w * 0.02,
          lineStyle: {
            color: "#fff",
            width: w * 0.005,
          },
        },
        splitLine: {
          distance: -(w * 0.1),
          length: w * 0.1,
          lineStyle: {
            color: "#fff",
            width: w * 0.005,
          },
        },
        // 数字刻度
        axisLabel: {
          color: "inherit",
          distance: w * 0.1,
          fontSize: w * 0.05,
        },

        // 数值
        detail: {
          valueAnimation: true,
          formatter: "{value}%",
          color: "inherit",
          fontSize: w * 0.09,
          // padding: [100, 0, 0, 0],
          offsetCenter: [0, w * 0.3],
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

const tick = async (param: TickParam) => {
  const cpuTimesResult = await getMachineCpuTimes();
  machineCpuTimes.value = cpuTimesResult.data;
  console.log("cpu times", cpuTimesResult, param);
  tickParam.value = param;
};
</script>
