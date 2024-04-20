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
  <div class="h-17" :class="[props.isTopSticky ? 'bg-[#EBEBEB]' : '', props.isCurrentChat ? 'bg-[#DEDEDE]!' : '']">前端艺术社
  </div>
</template>

<style scoped>

</style>
