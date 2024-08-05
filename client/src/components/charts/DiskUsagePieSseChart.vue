<template>
  <ResizableChart :option="option" @resize="onResize" />
</template>

<script lang="ts" setup>
import { computed, onBeforeMount, ref } from "vue";
import { ResizeParam } from "./ResizableChart.vue";
import { CjMachineDiskUsage, getMachineDiskUsage, watchMachineDiskUsage } from "../../apis/machine";

const idleColor = "#bbb";

const machineDiskUsage = ref<CjMachineDiskUsage[]>([]);
const resizeParam = ref<ResizeParam>();
const onResize = (param: ResizeParam) => {
  resizeParam.value = param;
};

const option = computed(() => {
  const w = resizeParam.value?.width ?? 0;

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

onBeforeMount(async () => {
  const diskUsageResult = await getMachineDiskUsage();
  machineDiskUsage.value = diskUsageResult.data;
  console.log('watchMachineDiskUsage sss');
  watchMachineDiskUsage(event => {
    const data = JSON.parse(event.data);
    console.log('watchMachineDiskUsage', event);
    machineDiskUsage.value = data;
  });
});
</script>
