import axios, {
  type InternalAxiosRequestConfig,
  type AxiosResponse,
  type AxiosInstance,
} from "axios";
import { useAuthStore } from "../stores/AuthStore";
import { isEmpty } from "lodash";
import { ElMessage } from "element-plus";

export const API_BASE_URL = process.env.NODE_ENV == "production" ? "" : "/api";

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
        ElMessage({
          message: error,
          grouping: true,
          type: "error",
        });
        return Promise.reject(error);
      }
    );
  }

  async get<R>(url: string): Promise<ApiResult<R>> {
    const response = await this.client.get(url);
    if (response.status == 200 && response.data.code == 0) {
      return response.data as ApiResult<R>;
    }
    throw new Error(`GET 失败：${url}`);
  }

  async post<P, R>(url: string, param: P): Promise<ApiResult<R>> {
    const response = await this.client.post(url, param);
    if (response.status == 200 && response.data.code == 0) {
      return response.data as ApiResult<R>;
    }
    throw new Error(`POST 失败：${url} ${param}`);
  }
}

export const api: ApiClient = new ApiClient();
