// 好友列表单个好友数据，对应 FriendItemDTO
export interface FRIEND {
    CreatedAt: string;
    Uid: string;
    Email: string;
    Avatar: string;
    Status: number;
}

// 渲染消息体结构
export type MSG_TYPE = "message" | "img" | "video" | "audio"

export interface MSG {
    from: string;
    to: string;
    type: MSG_TYPE;
    data: any;
}
