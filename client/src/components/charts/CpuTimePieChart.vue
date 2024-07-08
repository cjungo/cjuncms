<template>
  <VChart class="chart" :option="option" :autoresize="true" />
</template>

<script lang="ts" setup>
import VChart from "vue-echarts";
import { computed, onBeforeMount, ref } from "vue";
import { type CjMachineCPUTime, getMachineCpuTimes } from "../../apis/machine";
import { onBeforeRouteLeave, onBeforeRouteUpdate } from "vue-router";

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
  const cpu = machineCpuTimes.value ?? {
    system: 0,
    user: 0,
    idle: 0,
  };

  return {
    title: {
      text: "CPU 状态",
      left: "center",
    },
    tooltip: {
      trigger: "item",
      formatter: "{a} <br/>{b} : {c} ({d}%)",
    },
    legend: {
      orient: "vertical",
      left: "4%",
      data: ["CPU 系统占用", "CPU 用户占用", "CPU 空闲"],
    },
    series: [
      {
        name: "CPU",
        type: "pie",
        radius: ["10%", "40%"],
        center: ["50%", "50%"],
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
    ],
  };
});

const tick = async () => {
  const cpuTimesResult = await getMachineCpuTimes();
  machineCpuTimes.value = cpuTimesResult.data;
  console.log("cpu times", cpuTimesResult);
};
const tickId = ref(0);

onBeforeMount(async () => {
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
