import { apiGet } from "../utils/api";

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
