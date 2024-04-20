<script setup lang="ts">
import { computed, reactive, ref } from "vue";
import { useRouter } from "vue-router";
import { post } from "@/reuqest/request.ts";

const router = useRouter()

const activeKey = ref<"1" | "2">("1");

const btnText = computed(() => {
  return activeKey.value === "1" ? "登录" : "注册";
});

const validateMessages = {
  required: "${label} is required!",
  types: {
    email: "it is not a valid email!",
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

const clickMap = {
  "1": () => {
    handleLogIn();
  },
  "2": () => {

  }
};

const handleLogIn = async () => {
  console.log(formState.user);
  try {
    const res = await post("/api/client/signIn", {email: "22@qq.com", password: "1234567890"});
    console.log(res);
    await router.push(({name: "Chat"}));
  } catch (e) {
    console.error(e);
  }
};

const handleClick = () => {
  clickMap[activeKey.value]?.();
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
          <div class="flex-1 flex flex-col">
            <a-form
                class="h-full flex flex-col"
                :validate-messages="validateMessages"
                :model="formState"
            >
              <a-form-item :name="['user', 'email']" :rules="[{ required: true, type: 'email' }]">
                <a-input class="h-10" v-model:value="formState.user.email" placeholder="input email"/>
              </a-form-item>
              <a-form-item :name="['user', 'captcha']">
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
              <a-form-item :name="['user', 'password']">
                <a-input-password class="h-10"
                                  v-model:value="formState.user.password"
                                  placeholder="input password"
                />
              </a-form-item>
              <a-button class="h-10 mt-5" type="primary" @click="handleClick">{{ btnText }}</a-button>
            </a-form>
          </div>
        </div>
      </div>
    </div>
<!--    <img class="w-full h-full" src="https://w.wallhaven.cc/full/jx/wallhaven-jxl31y.png" alt="">-->
  </div>
</template>

<style scoped>

</style>
