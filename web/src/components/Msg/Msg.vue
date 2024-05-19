<script setup lang="ts">
import { MSG } from "@/types/data";

interface Props {
  data: MSG;
}

const props = defineProps<Props>();
</script>

<template>
  <div class="w-full ">
    <!-- from me-->
    <div class="flex justify-end" v-if="props.data.from === '1' && props.data.to !== '1'">
      <div class="bg-red p-1 max-w-60% text-wrap">{{ props.data.data }}</div>
      <img class="w-12 h-12 rd-1" src="https://avatars.githubusercontent.com/u/58922004?v=4" alt="">
    </div>

    <!--  to me-->
    <div v-else>
      别人发我的
    </div>
  </div>
</template>

<style scoped>

</style>
