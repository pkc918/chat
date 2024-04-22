<script setup lang="ts">

import { ref } from "vue";

const message = ref<string>("")
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
    聊天内容
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
    <a-textarea class="w-full! h-full! p-0! text-4 appearance-none! shadow-none! border-none! resize-none! outline-none! bg-transparent" v-model:value="message" placeholder="Basic usage" :rows="4" />
  </div>
</div>
</template>

<style scoped>

</style>
