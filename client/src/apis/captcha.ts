import { type ApiResult, api } from "../utils/api";

export type GetCaptchaMathResult = ApiResult<{
  id: string;
  image: string;
}>;

export const getCaptchaMath = async (): Promise<GetCaptchaMathResult> => {
  return await api.get("/captcha/math");
};
