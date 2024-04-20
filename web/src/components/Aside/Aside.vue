<script setup lang="ts">
import { onMounted, ref } from 'vue';
import { RouteRecordName, useRouter } from "vue-router";
import { useLocalRoute } from "@/store/useLoacalRoute.ts";

const router = useRouter()

interface routesItem{
  iconName: string;
  pathName: string;
}
const routesHashMapList: routesItem[] = [{
  iconName: "acorn-duotone",
  pathName:  "Chat"
}, {
  iconName: "android-logo-duotone",
  pathName: "Friend"
}, {
  iconName: "codesandbox-logo-duotone",
  pathName: "Collect"
}, {
  iconName: "aperture-duotone",
  pathName: "Comment"
}, {
  iconName: "battery-charging-vertical-duotone",
  pathName: "Battery"
}, {
  iconName: "cloud-duotone",
  pathName: "Weather"
}, {
  iconName: "calendar-dots-duotone",
  pathName: "Calendar"
}, {
  iconName: "dots-three-circle-duotone",
  pathName: "Setting"
}]

const activeRoute = useLocalRoute();
const activePath = ref<RouteRecordName>("")
const handleToRoute = (pathName: string) => {
  console.log(pathName)
  activePath.value = pathName
  router.push({name: pathName})
}

onMounted(() => {
  activePath.value = activeRoute.routeName;
})
</script>

<template>
<!-- unocss 识别icon，不然无法动态icon  -->
  <div v-if="false" class="i-ph-acorn-duotone i-ph-android-logo-duotone i-ph-codesandbox-logo-duotone i-ph-aperture-duotone i-ph-battery-charging-vertical-duotone i-ph-cloud-duotone i-ph-calendar-dots-duotone i-ph-dots-three-circle-duotone"></div>

  <div class="w-16 bg-[#CFD6DA] flex flex-col items-center justify-between">
    <div class="w-12 pt-17 flex flex-col items-center gap-8">
      <img class="w-12 h-12 rd-2" src="https://avatars.githubusercontent.com/u/58922004?v=4" alt="" />
      <div class="w-8 h-8" :class="[`i-ph-${r.iconName}`, `${r.pathName === activePath ? 'text-[#57bf6b]' : ''}`]" :key="r.iconName" v-for="r in routesHashMapList.slice(0,4)" @click="handleToRoute(r.pathName)" />
    </div>
    <div class="w-12 pb-8 flex flex-col items-center gap-8">
      <div class="w-8 h-8" :class="[`i-ph-${r.iconName}`, `${r.pathName === activePath ? 'text-[#57bf6b]' : ''}`]" :key="r.iconName" v-for="r in routesHashMapList.slice(4,routesHashMapList.length)" @click="handleToRoute(r.pathName)" />
    </div>
  </div>
</template>

<style scoped></style>
