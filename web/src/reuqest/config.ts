import { InterceptorMethodType } from "./types";
import { useUserStore } from "@/store/userStore.ts";

export const serverConfig = {
    baseURL: "",
    timeout: 5000,
    withCredentials: false,
    useTokenAuthorization: false,
};

// 请求拦截器
export const requestInterceptor: InterceptorMethodType = (config: any) => {
    console.log("request", config);
    config.headers.tenantId = "1";
    config.headers["authorization"] = useUserStore().token;
    return config;
};

// 请求错误拦截器
export const requestInterceptorError: InterceptorMethodType = (err: any) => {
    console.log("request error");
    return Promise.reject(err);
};

// 响应拦截器
export const responseInterceptor: InterceptorMethodType = (res: any) => {
    console.log("response", res);
    return res;
};

// 响应错误拦截器
export const responseInterceptorError: InterceptorMethodType = (error: any) => {
    console.log("response error");
    return Promise.reject(error);
};
