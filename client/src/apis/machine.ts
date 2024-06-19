import { type ApiResult, api } from "../utils/api";

export type MachineCpuInfo = {
  cpu: number;
  vendorId: string;
  family: string;
  model: string;
  stepping: number;
  physicalId: string;
  coreId: string;
  cores: number;
  modelName: string;
  mhz: number;
  cacheSize: number;
  flags: string[];
  microcode: string;
};

export type MachineCpuTime = {
  cpu: string;
  user: number;
  system: number;
  idle: number;
  nice: number;
  iowait: number;
  irq: number;
  softirq: number;
  steal: number;
  guest: number;
  guestNice: number;
};

export type MachineVirtualMemory = {
  total: number;
  available: number;
  used: number;
  usedPercent: number;
  free: number;
  active: number;
  inactive: number;
  wired: number;
  laundry: number;
  buffers: number;
  cached: number;
  writeBack: number;
  dirty: number;
  writeBackTmp: number;
  shared: number;
  slab: number;
  sreclaimable: number;
  sunreclaim: number;
  pageTables: number;
  swapCached: number;
  commitLimit: number;
  committedAS: number;
  highTotal: number;
  highFree: number;
  lowTotal: number;
  lowFree: number;
  swapTotal: number;
  swapFree: number;
  mapped: number;
  vmallocTotal: number;
  vmallocUsed: number;
  vmallocChunk: number;
  hugePagesTotal: number;
  hugePagesFree: number;
  hugePagesRsvd: number;
  hugePagesSurp: number;
  hugePageSize: number;
  anonHugePages: number;
};

export type GetMachineCpuInfoResult = ApiResult<MachineCpuInfo>;
export type GetMachineCpuTimesResult = ApiResult<[MachineCpuTime]>;
export type GetMachineVirtualMemoryResult = ApiResult<MachineVirtualMemory>;

export const getMachineCpuInfo = async (): Promise<GetMachineCpuInfoResult> => {
  return await api.get("/api/machine/cpu/info");
};

export const getMachineCpuTimes =
  async (): Promise<GetMachineCpuTimesResult> => {
    return await api.get("/api/machine/cpu/times");
  };

export const getMachineVirtualMemory =
  async (): Promise<GetMachineVirtualMemoryResult> => {
    return await api.get("/api/machine/virtual-memory");
  };
