<template>
  <TickableChart :option="option" @tick="tick" />
</template>

<script lang="ts" setup>
import { computed, ref } from "vue";
import { CjMachineDiskUsage, getMachineDiskUsage } from "../../apis/machine";
import { type TickParam } from "./TickableChart.vue";

const idleColor = "#bbb";

const machineDiskUsage = ref<CjMachineDiskUsage[]>([]);
const tickParam = ref<TickParam>();
const option = computed(() => {
  const w = tickParam.value?.width ?? 0;

  const names = machineDiskUsage.value
    .map((i) => [
      {
        name: `${i.path} 空闲`,
        itemStyle: {
          color: idleColor,
        },
      },
      {
        name: `${i.path} 使用`,
      },
    ])
    .flat();
  const pieCount = machineDiskUsage.value.length;
  const series = machineDiskUsage.value.map((i, index) => {
    const x = (1 / pieCount) * (index + 0.5);
    return {
      name: i.path,
      type: "pie",
      radius: [w * 0.02, w * 0.06],
      center: [w * x, w * 0.16],
      data: [
        {
          value: i.free,
          name: `${i.path} 空闲`,
          itemStyle: {
            color: idleColor,
          },
        },
        {
          value: i.used,
          name: `${i.path} 使用`,
        },
      ],
      animation: false,
      emphasis: {
        itemStyle: {
          shadowBlur: 10,
          shadowOffsetX: 0,
          shadowColor: "rgba(0, 0, 0, 0.5)",
        },
      },
    };
  });
  return {
    title: {
      text: "硬盘状态",
      left: "center",
      textStyle: {
        fontSize: w * 0.016,
      },
    },
    tooltip: {
      trigger: "item",
      formatter: "{a} <br/>{b} : {c} ({d}%)",
    },
    legend: {
      orient: "horizontal",
      top: w * 0.05,
      left: w * 0.04,
      itemWidth: w * 0.03,
      itemHeight: w * 0.014,
      data: names,
    },
    series: series,
    backgroundColor: "#fff",
  };
});

const tick = async (param: TickParam) => {
  const diskUsageResult = await getMachineDiskUsage();
  machineDiskUsage.value = diskUsageResult.data;
  console.log("disk usage", diskUsageResult);
  tickParam.value = param;
};
</script>

<style lang="scss" scoped></style>
