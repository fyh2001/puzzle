<script lang="ts" setup>
import TitleBar from "@/components/title-bar.vue";
import { ref, computed } from "vue";
import { md5 } from "js-md5";
import router from "@/routers/index";
import { salt } from "@/config/index";
import { useMessage } from "naive-ui";
import { userRequest } from "api/user";
import { useUserStore } from "@/store/user";
import { useI18n } from "vue-i18n";

const { t } = useI18n();

const userStore = useUserStore();

// 消息提示
const Message = useMessage();

// 表单实例
const formRef = ref(null);
// 表单数据
const formValue = ref({
  username: "",
  password: "",
});

// 验证规则
const regulation = {
  username: {
    label: t("login.byUsername.rules.username.label"),
    expression: /^[a-zA-Z0-9]{6,12}$/,
    message: t("login.byUsername.rules.username.message.pattern"),
  },
  password: {
    label: t("login.byUsername.rules.password.label"),
    expression: /^[a-zA-Z0-9]{6,12}$/,
    message: t("login.byUsername.rules.password.message.pattern"),
  },
};

// 登录按钮是否禁用
const loginButtonDisabled = computed(() => {
  return !(
    regulation.username.expression.test(formValue.value.username) &&
    regulation.password.expression.test(formValue.value.password)
  );
});

// 表单验证规则
const rules = {
  username: [
    {
      required: true,
      message: t("login.byUsername.rules.username.message.required"),
      trigger: "blur",
    },
    {
      pattern: regulation.username.expression,
      message: regulation.username.message,
      trigger: "blur",
    },
  ],
  password: [
    {
      required: true,
      message: t("login.byUsername.rules.password.message.required"),
      trigger: "blur",
    },
    {
      pattern: regulation.password.expression,
      message: regulation.password.message,
      trigger: "blur",
    },
  ],
};

/**
 * 登录
 */
const loginHandler = async () => {
  // 表单验证
  if (
    !(
      regulation.username.expression.test(formValue.value.username) &&
      regulation.password.expression.test(formValue.value.password)
    )
  ) {
    return Message.error(t("login.byUsername.rules.rulesError"));
  }

  // 登录
  const {
    data: { code, msg, data: userResp },
  } = await userRequest.login({
    username: formValue.value.username,
    password: md5(formValue.value.password + salt),
  });

  if (code === 200) {
    userStore.setUser(userResp.user); // 更新本地用户信息
    userStore.setToken(userResp.token); // 更新本地token

    // await saveUnUploadRecords(); // 保存未上传的记录

    Message.success(t("login.message.success")); // 提示登录成功
    return setTimeout(() => router.replace("/user"), 1000); // 1秒后返回上一页
  }

  Message.error(msg); // 提示登录失败
};
</script>

<template>
  <div class="flex justify-center items-center p-4">
    <n-el
      class="px-5 py-6 w-full rd-3 shadow translate-y-1/4"
      style="background: var(--login-panel-background-color)"
    >
      <title-bar class="mb-4" :title="t('login.title')" />
      <!-- 表单 -->
      <n-form ref="formRef" :label-width="80" :model="formValue" :rules="rules">
        <n-form-item path="username" style="--n-label-height: 0px">
          <n-input
            v-model:value="formValue.username"
            style="
              --n-height: 50px;
              --n-border-radius: 0.75rem;
              --n-border-hover: var(--primary-color);
              --n-border-focus: var(--primary-color);
            "
            :placeholder="t('login.byUsername.placeholder.username')"
            size="large"
          />
        </n-form-item>

        <n-form-item path="password" style="--n-label-height: 10px">
          <n-input
            v-model:value="formValue.password"
            type="password"
            show-password-on="click"
            style="
              --n-height: 50px;
              --n-border-radius: 0.75rem;
              --n-border-hover: var(--primary-color);
              --n-border-focus: var(--primary-color);
            "
            :placeholder="t('login.byUsername.placeholder.password')"
            size="large"
          />
        </n-form-item>
      </n-form>

      <!-- 按钮 -->
      <div class="">
        <n-button
          type="primary"
          class="w-full h-12 rounded-xl text-5 shadow"
          :disabled="loginButtonDisabled"
          @click="loginHandler"
          >{{ t("login.byUsername.button.login") }}</n-button
        >
      </div>

      <div class="flex justify-center items-center mt-4 text-gray gap-2">
        <div>{{ t("login.byUsername.other.loginByPhone") }}</div>
        <div>|</div>
        <div @click="$router.replace('register')">
          {{ t("login.byUsername.other.register") }}
        </div>
        <div>|</div>
        <div>{{ t("login.byUsername.other.forgotPassword") }}</div>
      </div>
    </n-el>
  </div>
</template>
@/store/user
