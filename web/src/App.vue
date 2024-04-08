<script setup lang="ts">
import Aside from "@/components/Aside/Aside.vue";
</script>

<template>
  <div class="w-full h-full flex">
    <Aside/>
    <main class="flex-1">
      <RouterView/>
    </main>
  </div>
</template>

<style scoped>

</style>
