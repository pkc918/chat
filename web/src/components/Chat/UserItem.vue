<script setup lang="ts">
interface Props{
  isTopSticky?: boolean
}
const props = withDefaults(defineProps<Props>(), {
  isTopSticky: false
})
</script>

<template>
  <div class="h-17" :class="[props.isTopSticky ? 'bg-[#EBEBEB]' : '']">前端艺术社</div>
</template>

<style scoped>

</style>
