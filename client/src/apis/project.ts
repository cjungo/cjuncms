import { apiDelete, apiPost } from "../utils/api";

export type CjProject = {
  id: number; //
  name: string; // 项目名
};

export type QueryProjectParam = {
  skip: number;
  take: number;
  plain: string;
};
export const queryProject = apiPost<QueryProjectParam, CjProject[]>(
  "/api/project/query"
);

export type AddProjectParam = Omit<CjProject, "id">;
export const addProject = apiPost<AddProjectParam, CjProject>(
  "/api/project/add"
);

export type EditProjectParam = CjProject;
export const editProject = apiPost<EditProjectParam, any>("/api/project/edit");

export type DropProjectParam = {
  id?: number;
  ids?: number[];
};
export const dropProject = apiDelete<DropProjectParam, any>(
  "/api/project/drop"
);
