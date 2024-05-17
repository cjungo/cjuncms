import axios, {
    type InternalAxiosRequestConfig,
    type AxiosInstance,
} from "axios";
import { useAuthStore } from "../stores/AuthStore";
import { isEmpty } from "lodash";

export const API_BASE_URL = process.env.NODE_ENV == 'production' ? '' : '/api';

export const api: AxiosInstance = axios.create({
    baseURL: API_BASE_URL,
    timeout: 40000,
});

api.interceptors.request.use(async (config: InternalAxiosRequestConfig<any>): Promise<InternalAxiosRequestConfig<any>> => {
    const auth = useAuthStore();
    if (!isEmpty(auth.token)) {
        config.headers['Authorization'] = `Bearer ${auth.token}`;
    }
    return config;
});