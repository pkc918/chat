<script setup lang="ts">

import { ref } from "vue";
import Msg from "@/components/Msg/Msg.vue";
import { MSG } from "@/types/data";
// import socket from "@/socket/socket.ts";

const message = ref<string>("")

const msgList = ref<MSG[]>([]);

const socket = new WebSocket(`ws://127.0.0.1:8080/api/client/v1/ws?id=${localStorage.getItem("id")}`);


socket.addEventListener("open", () => {
  console.log("WebSocket连接已建立");
});

socket.addEventListener("message", (event) => {
  msgList.value.push({
    from: "1",
    to: "2",
    type: "message",
    data: event.data
  });
  const message = event.data;
  console.log("收到消息：", message);
});

socket.addEventListener("close", () => {

  console.log("WebSocket连接已断开");
});

socket.addEventListener("error", (error) => {

  console.error("发生错误：", error);
});
const handleToSendMsg = () => {
  console.log("send Msg: ", message.value);
  socket.send(message.value);
};
</script>

<template>
<div class="w-full h-full flex flex-col justify-between">
  <header class="w-full min-h-15 max-h-20 border-b-1 border-b-[#E0E0E0] border-b-solid flex justify-between items-center px-5">
    <div>
      <h3>陪我去看海吧！(165)</h3>
    </div>
    <div class="i-ph-dots-three-circle-duotone text-8"></div>
  </header>
  <main class="flex-1 p-5">
    <Msg v-for="msg in msgList" :data="msg"/>
  </main>
  <div class="w-full h-65 flex flex-col border-t-1 border-t-[#E0E0E0] border-t-solid px-5">
    <div class="w-full flex justify-between py-5 text-8">
      <div class="flex gap-5">
        <div class="i-ph-bug-beetle-duotone hover:text-[#57bf6b]"></div>
        <div class="i-ph-tree-palm-duotone hover:text-[#57bf6b]"></div>
        <div class="i-ph-twitch-logo-duotone hover:text-[#57bf6b]"></div>
        <div class="i-ph-dots-three-circle-duotone hover:text-[#57bf6b]"></div>
      </div>
      <div class="flex gap-5">
        <div class="i-ph-film-reel-duotone hover:text-[#57bf6b]"></div>
        <div class="i-ph-phone-call-duotone hover:text-[#57bf6b]"></div>
      </div>
    </div>
    <a-textarea
        class="w-full! h-full! p-0! text-4 appearance-none! shadow-none! border-none! resize-none! outline-none! bg-transparent"
        v-model:value="message"
        placeholder="Basic usage"
        :rows="4"
        @keydown.enter="handleToSendMsg"
    />
  </div>
</div>
</template>

<style scoped>

</style>
