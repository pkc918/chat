import { get } from "@/reuqest/request.ts";

export const getFriend = <T>() => {
    return get<T>("/api/client/v1/friendList")
}
