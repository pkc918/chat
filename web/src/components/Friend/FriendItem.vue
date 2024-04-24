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
  <div class="flex gap-2 items-center p-2 px-5 h-15 bg-red">
    <img class="h-full" src="@/assets/vue.svg">
    前端艺术社
  </div>
</template>

<style scoped>

</style>
