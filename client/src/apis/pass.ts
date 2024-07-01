import { apiGet } from "../utils/api";

export type Pass = {};

export type QueryPassParam = {
  skip: number;
  take: number;
  plain: string;
};
export const queryPass = apiGet<QueryPassParam, Pass[]>("/api/pass/query");

