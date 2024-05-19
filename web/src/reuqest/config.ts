import { InterceptorMethodType } from "./types";

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
    config.headers["authorization"] = "eyJhbGciOiJIUzI1NiIsIkDt4-p7YaOxcZvX67406E6xnx8U0gnR5cCI6IkpXVCJ9.eyJpZCI6IjMzOGQ2NTliLWFjZWMtNDE4YS1hMGY3LTJhNGJlYzY4OGM2YiIsImVtYWlsIjoiMjJAcXEuY29tIiwicGFzc3dvcmQiOiIxMjM0NTY3ODkwIiwiZXhwIjoxNzE2MTgyODI5LCJpYXQiOjE3MTYwOTY0Mjl9.7D8tkYfwISY8zs"
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
    return res.data;
};

// 响应错误拦截器
export const responseInterceptorError: InterceptorMethodType = (error: any) => {
    console.log("response error");
    return Promise.reject(error);
};
