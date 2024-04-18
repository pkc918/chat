export interface InterceptorMethodType {
    <T>(query: T): T | Promise<T>;
}
