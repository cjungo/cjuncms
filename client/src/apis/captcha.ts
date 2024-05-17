import { api } from "../utils/api";

export type GetCaptchaMathResult = {
  code: number;
  data: {
    id: string;
    image: string;
  };
};

export const getCaptchaMath = async (): Promise<GetCaptchaMathResult> => {
  const response = await api.get("/captcha/math");
  return response.data
};
