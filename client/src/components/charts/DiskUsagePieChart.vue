<template>
  <VChart class="chart" :option="option" :autoresize="true" />
</template>

<script lang="ts" setup>
import VChart from "vue-echarts";
import { computed, onBeforeMount, ref } from "vue";
import { CjMachineDiskUsage, getMachineDiskUsage } from "../../apis/machine";
import { onBeforeRouteLeave, onBeforeRouteUpdate } from "vue-router";

const machineDiskUsage = ref<CjMachineDiskUsage[]>([]);

const option = computed(() => {
  const names = machineDiskUsage.value.map((i) => [`${i.path} 空闲`, `${i.path} 使用`]).flat();
  const pieCount = machineDiskUsage.value.length;
  const pieSize = 100.0 / pieCount * 2;
  const series = machineDiskUsage.value.map((i, index) => {
    const x = (100 / pieCount) * (index + 0.5);
    // console.log('x', x, pieSize);
    return {
      name: i.path,
      type: "pie",
      radius: ["10%", `${pieSize}%`],
      center: [`${x}%`, "50%"],
      data: [
        { value: i.free, name: `${i.path} 空闲` },
        { value: i.used, name: `${i.path} 使用` },
      ],
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
    },
    tooltip: {
      trigger: "item",
      formatter: "{a} <br/>{b} : {c} ({d}%)",
    },
    legend: {
      orient: "vertical",
      left: "4%",
      data: names,
    },
    series: series,
  };
});

const tick = async () => {
  const diskUsageResult = await getMachineDiskUsage();
  machineDiskUsage.value = diskUsageResult.data;
  console.log("disk usage", diskUsageResult);
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
