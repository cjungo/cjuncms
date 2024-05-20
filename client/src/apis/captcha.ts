import { type AxiosResponse } from "axios";
import { ApiResult, api } from "../utils/api";

export type GetCaptchaMathResult = ApiResult<{
  id: string;
  image: string;
}>;

export const getCaptchaMath = async (): Promise<
  AxiosResponse<GetCaptchaMathResult>
> => {
  return await api.get("/captcha/math");
};
