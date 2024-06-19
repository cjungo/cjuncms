import { type ApiResult, api } from "../utils/api";

export type MachineCpuInfo = ApiResult<{
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
  flags: [string];
  microcode: string;
}>;

export const getMachineCpuInfo = async (): Promise<MachineCpuInfo> => {
  return await api.get("/machine/cpu/info");
};
