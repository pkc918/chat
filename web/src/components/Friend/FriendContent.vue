<script setup lang="ts">
const handleToLink = () => {
  window.open('https://github.com/pkc918')
}
</script>

<template>
  <div class="w-full h-full flex items-center justify-center">
    <div class="flex flex-col gap-10 items-center">
      <img class="w-80 h-80 rd-full" src="https://avatars.githubusercontent.com/u/58922004?v=4" alt=""/>
      <div class="text-8">个性签名: 我追求那得不到的东西，我得到了我所没有追求的东西</div>
      <ol class="flex justify-around w-full">
        <li>
          <div class="i-ph-github-logo-duotone text-20" @click="handleToLink"></div>
        </li>
        <li>
          <div class="i-ph-github-logo-duotone text-20"></div>
        </li>
        <li>
          <div class="i-ph-github-logo-duotone text-20"></div>
        </li>
      </ol>
    </div>
  </div>
</template>

<style scoped>

</style>
