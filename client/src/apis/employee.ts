import { apiDelete, apiGet, apiPost, apiPut } from "../utils/api";

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

export type AddEmplyeeParam = Omit<Employee, "id">;
export const addEmplyee = apiPut<AddEmplyeeParam, Employee>(
  "/api/employee/add"
);

export type EditEmployeeParam = Employee;
export const editEmpoyee = apiPost<EditEmployeeParam, any>("api/employee/edit");

export type DropEmployeeParam = {
  id?: number;
  ids?: number[];
};
export const dropEmployee = apiDelete<DropEmployeeParam, any>(
  "/api/employee/drop"
);
