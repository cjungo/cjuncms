import { apiGet } from "../utils/api";

export type CaptchaMath = {
  id: string;
  image: string;
};

export const getCaptchaMath = apiGet<any, CaptchaMath>("/captcha/math");
