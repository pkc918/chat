<script setup lang="ts">
import { FRIEND } from "@/types/data";

interface Props{
  isTopSticky?: boolean,
  isCurrentChat?: boolean,
  friendInfo: FRIEND
}
const props = withDefaults(defineProps<Props>(), {
  isTopSticky: false,
  isCurrentChat: false,
})


</script>

<template>
  <div class="flex gap-2 items-center p-2 px-5 h-15 bg-red">
    <img class="h-full" :src="props.friendInfo.Avatar" alt="">
    {{ props.friendInfo.Uid }}
  </div>
</template>

<style scoped>

</style>
