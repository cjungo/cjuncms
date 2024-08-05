import axios, {
  type InternalAxiosRequestConfig,
  type AxiosResponse,
  type AxiosInstance,
} from "axios";
import { useAuthStore } from "../stores/AuthStore";
import { isEmpty } from "lodash";
import { tipError } from "./tip";
import router from "../router";

export const API_BASE_URL =
  process.env.NODE_ENV == "production" ? "" : "/proxy";

export type ApiResult<T> = {
  code: number;
  data: T;
};

export class ApiClient {
  private client: AxiosInstance;

  constructor() {
    this.client = axios.create({
      baseURL: API_BASE_URL,
      timeout: 40000,
    });
    this.client.interceptors.request.use(
      async (
        config: InternalAxiosRequestConfig<any>
      ): Promise<InternalAxiosRequestConfig<any>> => {
        const auth = useAuthStore();
        if (!isEmpty(auth.token)) {
          config.headers["Authorization"] = `${auth.token}`;
        }
        return config;
      }
    );
    this.client.interceptors.response.use(
      async (response: AxiosResponse<any>): Promise<AxiosResponse<any>> => {
        return response;
      },
      async (error: any): Promise<any> => {
        const result = error.response.data;
        console.log("result", result);
        if (result.code == -4) {
          const auth = useAuthStore();
          auth.token = "";
          await router.replace("/login");
        }
        tipError.next(error);
        return Promise.reject(error);
      }
    );
  }

  async get<P, R>(url: string, param?: P): Promise<ApiResult<R>> {
    const response = await this.client.get(url, {
      data: param,
    });
    if (response.status == 200 && response.data.code == 0) {
      return response.data as ApiResult<R>;
    }
    throw new Error(`GET 失败：${url}`);
  }

  async put<P, R>(url: string, param: P): Promise<ApiResult<R>> {
    const response = await this.client.put(url, param);
    if (response.status == 200 && response.data.code == 0) {
      return response.data as ApiResult<R>;
    }
    throw new Error(`PUT 失败：${url} ${param}`);
  }

  async post<P, R>(url: string, param: P): Promise<ApiResult<R>> {
    const response = await this.client.post(url, param);
    if (response.status == 200 && response.data.code == 0) {
      return response.data as ApiResult<R>;
    }
    throw new Error(`POST 失败：${url} ${param}`);
  }

  async delete<P, R>(url: string, param?: P): Promise<ApiResult<R>> {
    const response = await this.client.delete(url, {
      data: param,
    });
    if (response.status == 200 && response.data.code == 0) {
      return response.data as ApiResult<R>;
    }
    throw new Error(`DELETE 失败：${url}`);
  }
}

export const api: ApiClient = new ApiClient();

export const apiGet = <P, R>(url: string) => {
  return async (param?: P): Promise<ApiResult<R>> => {
    return await api.get<P, R>(url, param);
  };
};

export const apiPost = <P, R>(url: string) => {
  return async (param: P): Promise<ApiResult<R>> => {
    return await api.post<P, R>(url, param);
  };
};

export const apiPut = <P, R>(url: string) => {
  return async (param: P): Promise<ApiResult<R>> => {
    return await api.put<P, R>(url, param);
  };
};

export const apiDelete = <P, R>(url: string) => {
  return async (param?: P): Promise<ApiResult<R>> => {
    return await api.delete<P, R>(url, param);
  };
};

export const apiSse = (url: string, eventType: string = "message") => {
  return (eventHandler: (e: MessageEvent<string>) => void) => {
    const eventSource = new EventSource(url, {
      withCredentials: true,
    });
    eventSource.addEventListener(eventType, eventHandler);
    return eventSource;
  };
};
