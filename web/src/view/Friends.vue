<script lang="ts" setup>
import Layout from "@/components/Layout/Layout.vue";
import FriendItem from "@/components/Friend/FriendItem.vue";
import FriendContent from "@/components/Friend/FriendContent.vue";
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
            <FriendItem/>
          </div>
        </template>
        <template #content>
          <FriendContent/>
        </template>
      </Layout>
    </div>
</template>

<style scoped></style>
