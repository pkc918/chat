<script setup lang="ts">

import Layout from "@/components/Layout/Layout.vue";
import UserItem from "@/components/Chat/UserItem.vue";
</script>

<template>
  <Layout>
    <template #left>
      <div class="flex items-center justify-center gap-3 h-15">
        <input class="w-52 h-6.5" />
        <div class="w-6.5 h-6.5 bg-blue"></div>
      </div>
      <div>
        <UserItem></UserItem>
        <UserItem></UserItem>
        <UserItem></UserItem>
        <UserItem></UserItem>
        <UserItem></UserItem>
        <UserItem></UserItem>
      </div>

    </template>
    <template #content>content</template>
  </Layout>
</template>

<style scoped>

</style>
