import { apiPost } from "../utils/api";

export type Shell = {

};

export type QueryShellParam = {

};

export const queryShell = apiPost<QueryShellParam, Shell[]>("/api/shell/query");
