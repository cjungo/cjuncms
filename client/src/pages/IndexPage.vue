<template>
  <div class="index-page">
    <ElAutoResizer>
      <template #default>
        <v-chart class="chart" :option="option" />
      </template>
    </ElAutoResizer>
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
import { ref, provide, onBeforeMount } from "vue";
import { getMachineCpuInfo } from '../apis/machine'

use([
  CanvasRenderer,
  PieChart,
  TitleComponent,
  TooltipComponent,
  LegendComponent,
]);

// provide(THEME_KEY, "dark");

const option = ref({
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
    left: "left",
    data: ["CPU 使用", "CPU 空闲", "内存占用", "内存空闲"],
  },
  series: [
    {
      name: "CPU",
      type: "pie",
      radius: ["20%", "40%"],
      center: ["30%", "50%"],
      data: [
        { value: 200, name: "CPU 使用" },
        { value: 800, name: "CPU 空闲" },
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
      radius: ["20%", "40%"],
      center: ["70%", "50%"],
      data: [
        { value: 20, name: "内存占用" },
        { value: 80, name: "内存空闲" },
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
});

onBeforeMount(async () => {
  const response = await getMachineCpuInfo();
  
});

</script>

<style lang="scss" scoped>
.index-page {
  flex-grow: 1;
}
</style>
