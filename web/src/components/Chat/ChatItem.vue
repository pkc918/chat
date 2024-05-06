<script setup lang="ts">
interface Props{
  isTopSticky?: boolean,
  isCurrentChat?: boolean
}
const props = withDefaults(defineProps<Props>(), {
  isTopSticky: false,
  isCurrentChat: false
})


</script>

<template>
  <div class="h-17 cursor-pointer flex gap-2 items-center px-5" :class="[props.isTopSticky ? 'bg-[#EBEBEB]' : '', props.isCurrentChat ? 'bg-[#DEDEDE]!' : '']">
    <img class="w-12 h-12 rd-2" src="https://avatars.githubusercontent.com/u/58922004?v=4" alt=""/>

    <div>
      <div class="text-black">前端艺术社</div>
      <div class="text-3 text-gray">信息</div>
    </div>
  </div>
</template>

<style scoped>

</style>
