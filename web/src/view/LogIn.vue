<script setup lang="ts">
import { reactive, ref } from "vue";

const activeKey = ref("2");

const validateMessages = {
  required: "${label} is required!",
  types: {
    email: "it is not a valid email!",
    number: "${label} is not a valid number!",
  },
  number: {
    range: "${label} must be between ${min} and ${max}",
  },
};

interface Query {
  email: string;
  password: string;
  captcha?: string;
}

const formState = reactive<{ user: Query }>({
  user: {
    email: "",
    password: "",
    captcha: ""
  },
});

const handleLogIn = () => {
  console.log(formState.user);
};
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
        <div>
          <a-tabs v-model:activeKey="activeKey" type="card">
            <a-tab-pane key="1" tab="登录"></a-tab-pane>
            <a-tab-pane key="2" tab="注册"></a-tab-pane>
          </a-tabs>
        </div>
        <div class="flex-1 flex flex-col justify-between gap-5">
          <div class="flex-1 flex flex-col gap-5">
            <a-form
                class="h-full flex flex-col gap-5"
                :validate-messages="validateMessages"
                :model="formState"
            >
              <a-form-item :name="['user', 'email']" :rules="[{ required: true, type: 'email' }]">
                <a-input class="h-10" v-model:value="formState.user.email" placeholder="input email"/>
              </a-form-item>

              <a-form-item :name="['user', 'captcha']" :rules="[{ required: true, type: 'number', min: 6, max: 6 }]">
                <a-input-search
                    v-model:value="formState.user.captcha"
                    placeholder="input captcha"
                    size="large"
                >
                  <template #enterButton>
                    <a-button>获取验证码</a-button>
                  </template>
                </a-input-search>
              </a-form-item>

              <a-form-item :name="['user', 'email']" :rules="[{ required: true, type: 'email' }]">
                <a-input-password class="h-10"
                                  v-model:value="formState.user.password"
                                  placeholder="input password"
                />
              </a-form-item>
            </a-form>
          </div>
          <a-button class="h-10" type="primary" @click="handleLogIn">登录</a-button>
        </div>
      </div>
    </div>
<!--    <img class="w-full h-full" src="https://w.wallhaven.cc/full/jx/wallhaven-jxl31y.png" alt="">-->
  </div>
</template>

<style scoped>

</style>
