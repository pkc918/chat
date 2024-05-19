const socket = new WebSocket("ws://127.0.0.1:8080/api/client/v1/ws")


socket.addEventListener('open', () => {

    console.log('WebSocket连接已建立');
});

socket.addEventListener('message', (event) => {

    const message = event.data;
    console.log('收到消息：', message);
});

socket.addEventListener('close', () => {

    console.log('WebSocket连接已断开');
});

socket.addEventListener('error', (error) => {

    console.error('发生错误：', error);
});

export default socket
