import { get } from "@/reuqest/request.ts";

const getFriend = () => {
    return get("/client/v1/friendList")
}
