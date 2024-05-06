<script setup lang="ts">
const handleToLink = () => {
  window.open('https://github.com/pkc918')
}
</script>

<template>
  <div class="w-full h-full flex justify-center bg-white p-20">
    <div class="flex flex-col gap-10 items-start w-full">
      <div class="flex">
        <img class="w-20 h-20 rd-2" src="https://avatars.githubusercontent.com/u/58922004?v=4" alt=""/>
        <div>
          <ol>
            <li>陪我去看海</li>
            <li>Uid: 12345f-dskalf-jdls</li>
          </ol>
        </div>
      </div>
      <div class="text-5">个性签名: 我追求那得不到的东西，我得到了我所没有追求的东西</div>
      <ol class="flex flex-col w-full">
        <li class="flex items-center cursor-pointer" @click="handleToLink">
          <div class="i-ph-github-logo-duotone text-5"></div>
          <span class="text-5">: pkc918</span>
        </li>
      </ol>
    </div>
  </div>
</template>

<style scoped>

</style>
