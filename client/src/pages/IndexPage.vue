<template>
  <div class="index-page">
    <div class="machine-chart">
      <div class="machine-chart-box">
        <CpuTimeStageSpeedGaugeSseChart />
      </div>
      <div class="machine-chart-box">
        <VirtualMemorySpeedGaugeSseChart />
      </div>
    </div>
    <div class="machine-chart-row">
      <DiskUsagePieSseChart />
    </div>
    <div class="machine-chart-row">
      <CpuTimeLineSseChart />
    </div>
  </div>
</template>

<script setup lang="ts">
import { getMachineCpuInfo, type MachineCpuInfo } from "../apis/machine";
import { onBeforeMount, ref } from "vue";

const machineCpuInfo = ref<MachineCpuInfo>();

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
  background-color: #f6f8f9;
}
.machine-chart {
  display: flex;
  flex-direction: row;
  padding: 0 0.5vw;
}
.machine-chart-box {
  width: 14vw;
  height: 14vw;
  margin: 1vw 0.5vw;
  background-color: #fff;
}
.machine-chart-row {
  display: flex;
  flex-direction: row;
  margin: 0 1vw 1vw 1vw;
  min-height: 24vw;
}
</style>
