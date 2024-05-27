<script setup lang="ts">
import { MSG } from "@/types/data";

interface Props {
  data: MSG;
}

const id = localStorage.getItem("id");
const props = defineProps<Props>();
console.log("对比", id, props.data, id === props.data.from);

</script>

<template>
  <div class="w-full ">
    <!-- from me-->
    <div class="flex justify-end" v-if="props.data.from === id && props.data.to !== id">
      <div class="bg-red rd-2 p-1 px-2 max-w-60% text-wrap mx-2">{{ props.data.data }}</div>
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
