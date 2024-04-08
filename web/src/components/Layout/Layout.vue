<script setup lang="ts">

</script>

<template>
<div class="w-full h-full flex">
  <div class="w-70 h-full bg-[#F7F7F7] border-r-1 border-r-[#E0E0E0] border-r-solid">
    <slot name="left"></slot>
  </div>
  <div class="flex-1 h-full bg-[#F3F3F3]">
    <slot name="content"></slot>
  </div>
</div>
</template>

<style scoped>

</style>
