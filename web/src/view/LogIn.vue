<script setup lang="ts">
import { ref } from "vue";

const activeKey = ref("2");
</script>

<template>
  <div
      class="w-full h-full flex flex-col items-center justify-center login-bg">
    <h1>WeChat Plus</h1>
    <div class="w-200 h-100 mt-5 rd-2 bg-[rgba(255,255,255,.5)] flex">
      <div class="w-50% p-5">
        <div class="i-ph-alien-fill text-[#57bf6b] w-100% h-100%"></div>
      </div>
      <div class="w-50% p-10 flex flex-col">
        <div class="">
          <a-tabs v-model:activeKey="activeKey" type="card">
            <a-tab-pane key="1" tab="登录"></a-tab-pane>
            <a-tab-pane key="2" tab="注册"></a-tab-pane>
          </a-tabs>
        </div>
        <div class="flex-1 flex flex-col justify-between gap-5">
          <div class="flex flex-col gap-5">
            <a-input class="h-10" placeholder="input email"/>
            <a-input-password class="h-10"
                              placeholder="input password"
            />
          </div>
          <a-button class="h-10" type="primary">登录</a-button>
        </div>
      </div>
    </div>
<!--    <img class="w-full h-full" src="https://w.wallhaven.cc/full/jx/wallhaven-jxl31y.png" alt="">-->
  </div>
</template>

<style scoped>

</style>
