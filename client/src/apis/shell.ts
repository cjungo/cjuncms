import { apiGet } from "../utils/api";

export type Shell = {

};

export type QueryShellParam = {

};

export const queryShell = apiGet<QueryShellParam, Shell[]>("/api/shell/query");
