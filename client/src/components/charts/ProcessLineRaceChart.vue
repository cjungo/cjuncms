<template>
    <TickableChart :option="option" @tick="tick" />
</template>

<script lang="ts" setup>
import { computed, ref } from 'vue';
import { TickParam } from './TickableChart.vue';
import { type CjMachineProcess, getMachineProcesses } from '../../apis/machine';

const machineProcesses = ref<CjMachineProcess[]>([]);

const tickParam = ref<TickParam>();
const option = computed(() => {

});

const tick = async (param: TickParam) => {
  const result = await getMachineProcesses();
  machineProcesses.value = result.data;
  console.log("cpu times", result, param);
  tickParam.value = param;
};
</script>