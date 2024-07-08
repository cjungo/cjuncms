import { apiGet } from "../utils/api";

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

export type CjMachineCPUTime = {
  id: number; //
  cpu: string; //
  user: number; // 用户时间
  system: number; // 系统时间
  idle: number; // 空闲时间
  nice: number; //
  iowait: number; // IO等待
  irq: number; // 硬中断
  softirq: number; // 软中断
  steal: number; //
  guest: number; //
  guest_nice: number; //
  create_at: string; // 记录时间
};

export type CjMachineDiskUsage = {
  id: number; //
  path: string; //
  fstype: string; //
  total: number; //
  free: number; //
  used: number; //
  used_percent: number; //
  inodes_total: number; //
  inodes_used: number; //
  inodes_free: number; //
  inodes_used_percent: number; //
  create_at: string; // 记录时间
};

export type CjMachineProcess = {
  id: number; //
  pid: number; //
  name: string; //
  username: string; //
  cmdline: string; //
  workdir: string; //
  cpu_percent: number; //
  mem_percent: number; //
  create_at: string; // 记录时间
};

export type CjMachineVirtualMemory = {
  id: number; //
  total: number; // 全部
  available: number; //
  used: number; // 已用
  used_percent: number; //
  free: number; // 空闲
  active: number; //
  inactive: number; //
  wired: number; //
  laundry: number; //
  buffers: number; //
  cached: number; //
  write_back: number; //
  dirty: number; //
  write_back_tmp: number; //
  shared: number; //
  slab: number; //
  sreclaimable: number; //
  sunreclaim: number; //
  page_tables: number; //
  swap_cached: number; //
  commit_limit: number; //
  committed_as: number; //
  high_total: number; //
  high_free: number; //
  low_total: number; //
  low_free: number; //
  swap_total: number; //
  swap_free: number; //
  mapped: number; //
  vmalloc_total: number; //
  vmalloc_used: number; //
  vmalloc_chunk: number; //
  huge_pages_total: number; //
  huge_pages_free: number; //
  huge_pages_rsvd: number; //
  huge_pages_surp: number; //
  huge_page_size: number; //
  anon_huge_pages: number; //
  create_at: string; // 记录时间
};

export const getMachineCpuInfo = apiGet<any, MachineCpuInfo>(
  "/api/machine/cpu/info"
);

export const getMachineCpuTimes = apiGet<any, CjMachineCPUTime>(
  "/api/machine/cpu/times"
);

export type ListMachineCpuTimesParam = {
  startAt: string;
  endAt: string;
};
export const listMachineCpuTimes = apiGet<
  ListMachineCpuTimesParam,
  CjMachineCPUTime[]
>("/api/machine/cpu/times/list");

export const getMachineDiskUsage = apiGet<any, CjMachineDiskUsage[]>(
  "/api/machine/disk"
);

export const getMachineVirtualMemory = apiGet<any, CjMachineVirtualMemory>(
  "/api/machine/virtual-memory"
);

export type ListMachineVirtualMemoryParam = {
  startAt: string;
  endAt: string;
};
export const listMachineVirtualMemory = apiGet<
  ListMachineVirtualMemoryParam,
  CjMachineVirtualMemory[]
>("/api/machine/virtual-memory/list");

export const getMachineProcesses = apiGet<any, CjMachineProcess[]>(
  "/api/machine/processes"
);

export type ListMachineProcessesParam = {
  startAt: string;
  endAt: string;
};
export const listMachineProcesses = apiGet<
  ListMachineProcessesParam,
  CjMachineProcess[]
>("/api/machine/processes/list");
