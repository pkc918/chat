<script lang="ts" setup>
import Layout from "@/components/Layout/Layout.vue";
import FriendItem from "@/components/Friend/FriendItem.vue";
import FriendContent from "@/components/Friend/FriendContent.vue";
import { getFriend } from "@/api/friend.ts";
import { onMounted, ref } from "vue";
import { FRIEND } from "@/types/data";


const friendList = ref<FRIEND[]>([]);
const getFriendsForAuth = async () => {
  const res = await getFriend<FRIEND[]>();
  friendList.value = res.data;
  console.log("返回", res);
};
onMounted(() => {
  getFriendsForAuth();
});
</script>

<template>
    <div class="w-full h-full">
      <Layout>
        <template #left>
          <div class="flex items-center justify-center gap-3 h-15">
            <input class="w-52 h-6.5" />
            <div class="w-6.5 h-6.5 bg-blue"></div>
          </div>
          <div>
            <FriendItem v-for="friend in friendList" :key="friend.Uid" :friend-info="friend" />
          </div>
        </template>
        <template #content>
          <FriendContent/>
        </template>
      </Layout>
    </div>
</template>

<style scoped></style>
