<script setup lang="ts">
import { ref } from "vue";
import Layout from "@/components/Layout/Layout.vue";
import UserItem from "@/components/Chat/ChatItem.vue";
import ChatContent from "@/components/Chat/ChatContent.vue";


const topStickyGroup = ref<string[]>(["1"]);
const currentChatId = ref<string>("");

const handleToChat = (id: string) => {
  currentChatId.value = id;
};
</script>

<template>
  <Layout>
    <template #left>
      <div class="flex items-center justify-center gap-3 h-15">
        <input class="w-52 h-6.5" />
        <div class="w-6 h-6 cursor-pointer i-ph-first-aid-bold"></div>
      </div>
      <div>
        <UserItem :is-top-sticky="topStickyGroup.includes(item.toString())"
                  :is-current-chat="currentChatId === item.toString()" v-for="item in 6" :key="item"
                  @click="handleToChat(item.toString())"></UserItem>
      </div>

    </template>
    <template #content>
      <ChatContent/>
    </template>
  </Layout>
</template>

<style scoped>

</style>
