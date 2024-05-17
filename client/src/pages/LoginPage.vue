<template>
  <div class="login-page">
    <form class="login-form" @submit.prevent="onSubmit">
      <div class="login-title">
        <img alt="logo" src="/vite.svg" />
      </div>
      <div class="login-input">
        <label>用户</label>
        <input v-model="param.username" />
      </div>
      <div class="login-input">
        <label>密码</label>
        <input v-model="param.password" type="password" />
      </div>
      <div class="login-row">
        <label>验证码</label>
        <input v-model="param.captchaAnswer" />
        <img
          class="captcha"
          alt="captcha"
          :src="captchaSrc"
          @click="onClickCaptcha"
        />
      </div>
      <div class="login-row">
        <button type="submit">登录</button>
      </div>
    </form>
  </div>
</template>

<script setup lang="ts">
import { reactive, ref } from "vue";
import { useRouter } from "vue-router";
import { type LoginParam, login } from "../apis/login";
import { getCaptchaMath } from "../apis/captcha";
// import { useAuthStore } from "../stores/AuthStore";

const param = reactive<LoginParam>({
  username: "admin",
  password: "123456",
  captchaId: "",
  captchaAnswer: "",
});
const captchaSrc = ref("");

// const auth = useAuthStore();
const router = useRouter();

const onClickCaptcha = async () => {
  const captchaResult = await getCaptchaMath();
  captchaSrc.value = captchaResult.data.image;
  param.captchaId = captchaResult.data.id;
};

const onSubmit = async () => {
  const loginResult = await login(param);
  console.log("login result", loginResult);
  //TODO  auth
  router.push("/");
};
</script>

<style scoped lang="scss">
.login-page {
  display: flex;
  justify-content: center;
  align-items: center;
  width: 100%;
  height: 100%;
  background-size: cover;
  background-image: url("../assets/bg.png");
}

.login-form {
  display: flex;
  overflow: hidden;
  justify-content: center;
  align-items: stretch;
  flex-direction: column;
  border: 1px solid #d4d4d4;
  padding: 0.4em;
  flex-shrink: 1;
}

.login-title {
  display: flex;
  justify-content: center;
  align-items: center;
  border-bottom: 1px solid #eee;
  padding-bottom: 1em;
  margin: 2em 1em 1em 1em;

  & > img {
    width: 34%;
    user-select: none;
  }
}

.login-input {
  display: flex;
  align-items: center;
  flex-grow: 1;
  margin: 0.4em 0;

  & > label {
    min-width: 4em;
    text-align: left;
    text-indent: 1em;
    user-select: none;
  }

  & > input {
    flex-grow: 1;
    outline: none;
    border: 1px solid #f4f4f4;
    border-radius: 0.4em;
    padding: 0.4em;
    margin-left: 0.4em;
  }
}

.login-row {
  display: flex;
  align-items: center;
  margin: 0.4em 0;

  & > button {
    flex-grow: 1;
    padding: 0.4em;
    border: none;
    color: white;
    background-color: #49f;
    user-select: none;
  }

  & > label {
    min-width: 4em;
    text-align: left;
    text-indent: 1em;
    user-select: none;
  }

  & > input {
    flex-grow: 1;
    outline: none;
    border: 1px solid #f4f4f4;
    border-radius: 0.4em;
    padding: 0.4em;
    margin-left: 0.4em;
  }

  & > img.captcha {
    flex-grow: 1;
    margin-left: 1em;
    cursor: pointer;
    user-select: none;
  }
}
</style>
