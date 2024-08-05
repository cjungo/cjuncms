<template>
  <ResizableChart :option="option" @resize="onResize" />
</template>

<script lang="ts" setup>
import { computed, onBeforeMount, ref } from "vue";
import {
  getMachineCpuTimes,
  watchMachineCpuTimes,
  type CjMachineCPUTime,
} from "../../apis/machine";
import { ResizeParam } from "./ResizableChart.vue";
import { onBeforeRouteLeave } from "vue-router";

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

const resizeParam = ref<ResizeParam>();
const onResize = (param: ResizeParam) => {
  resizeParam.value = param;
};
const option = computed(() => {
  const w = resizeParam.value?.width ?? 0;

  const used = machineCpuTimes.value.user + machineCpuTimes.value.system;
  const total = machineCpuTimes.value.idle + used;
  const percent = (used / total) * 100;
  return {
    title: {
      text: "CPU",
      left: "center",
      textStyle: {
        fontSize: w * 0.08,
      },
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

onBeforeMount(async () => {
  const cpuTimesResult = await getMachineCpuTimes();
  machineCpuTimes.value = cpuTimesResult.data;
  console.log("cpu times", cpuTimesResult);
  watchMachineCpuTimes((event) => {
    const data = JSON.parse(event.data);
    machineCpuTimes.value = data;
  });
});

onBeforeRouteLeave(() => {

});
</script>
