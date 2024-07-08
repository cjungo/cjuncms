<template>
  <div class="index-page">
    <div class="machine-chart">
      <div class="machine-chart-box">
        <CpuTimePieChart />
      </div>
      <div class="machine-chart-box">
        <VirtualMemoryPieChart />
      </div>
    </div>
    <DiskUsagePieChart/>
  </div>
</template>

<script setup lang="ts">
import CpuTimePieChart from "../components/charts/CpuTimePieChart.vue";

import { getMachineCpuInfo, type MachineCpuInfo } from "../apis/machine";
import { onBeforeMount, ref } from "vue";
import DiskUsagePieChart from "../components/charts/DiskUsagePieChart.vue";

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

onBeforeMount(async () => {
  const cpuInfoResult = await getMachineCpuInfo();
  machineCpuInfo.value = cpuInfoResult.data;
  console.log("cpu", cpuInfoResult);
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
  display: flex;
  flex-direction: column;
  width: 100%;
  flex-grow: 1;
  overflow: auto;
}
.machine-chart {
  display: flex;
  flex-direction: row;
}
.machine-chart-box {
  width: 40vw;
  height: 40vw;
}
</style>
