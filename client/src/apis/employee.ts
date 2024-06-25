import { apiGet } from "../utils/api";

export type Employee = {
  id: number;
  jobnumber: string;
  username: string;
  nickname: string;
  fullname: string;
  birthday: string;
  avatar_path?: string;
  is_removed: number;
};

export type QueryEmployeeParam = {
  skip: number;
  take: number;
  plain: string;
};

export const queryEmployee = apiGet<QueryEmployeeParam, Employee[]>(
  "/api/employee/query"
);
