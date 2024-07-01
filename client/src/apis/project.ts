import { apiDelete, apiGet, apiPost, apiPut } from "../utils/api";

export type Project = {
  id: number;
  name: string;
};

export type QueryProjectParam = {
  skip: number;
  take: number;
  plain: string;
};
export const queryProject = apiGet<QueryProjectParam, Project[]>(
  "/api/project/query"
);

export type AddProjectParam = Omit<Project, "id">;
export const addProject = apiPut<AddProjectParam, Project>("/api/project/add");

export type EditProjectParam = Project;
export const editProject = apiPost<EditProjectParam, any>("/api/project/edit");

export type DropProjectParam = {
  id?: number;
  ids?: number[];
};
export const dropProject = apiDelete<DropProjectParam, any>("/api/project/drop");
