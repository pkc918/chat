<script setup lang="ts">
import Aside from "@component/Aside/Aside.vue";
import Layout from "@component/Layout/Layout.vue";
</script>

<template>
  <div class="w-full h-full flex">
    <Aside/>
    <main class="flex-1 bg-red">
      <RouterView/>
      <Layout/>
    </main>
  </div>
</template>

<style scoped>

</style>
