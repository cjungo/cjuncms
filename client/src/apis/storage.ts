import { apiPost, type ApiResult } from "../utils/api";

const uploadAvatarForm = apiPost<FormData, string>("/upload/avatar");

export const uploadAvatar = async (
  file: Blob,
  name: string
): Promise<ApiResult<string>> => {
  const param = new FormData();
  param.append("file", file, name);
  return await uploadAvatarForm(param);
};
