import axios, { AxiosInstance } from "axios";

import {
    requestInterceptor,
    requestInterceptorError,
    responseInterceptor,
    responseInterceptorError,
    serverConfig,
} from "./config.ts";

const serviceAxios: AxiosInstance = axios.create({
    baseURL: serverConfig.baseURL,
    timeout: serverConfig.timeout,
    withCredentials: serverConfig.withCredentials,
});

serviceAxios.interceptors.request.use(
    (config) => {
        return requestInterceptor<typeof config>(config);
    },
    (error) => {
        return requestInterceptorError<typeof error>(error);
    }
);
serviceAxios.interceptors.response.use(
    (res) => {
        return responseInterceptor<typeof res>(res);
    },
    (error) => {
        return responseInterceptorError<typeof error>(error);
    }
);

const post = <T>(url: string, data?: any) => {
    return serviceAxios<T>({
        url,
        data,
        method: "post",
    });
};

const get = <T>(url: string, params?: any) => {
    return serviceAxios<T>({
        url,
        params,
        method: "get",
    });
};

export { serviceAxios, post, get };
