<script lang="ts" setup>
import { ref, computed } from "vue";
import TitleBar from "@/components/title-bar.vue";
import router from "@/routers/index";
import { md5 } from "js-md5";
import { salt } from "@/config/index";
import { useMessage } from "naive-ui";
import { userRequest } from "api/user";
import { useI18n } from "vue-i18n";

const { t } = useI18n();

// 消息提示
const Message = useMessage();

// 表单实例
const formRef = ref(null);
// 表单数据
const formValue = ref({
  username: "",
  nickname: "",
  password: "",
});

// 验证规则
const regulation = {
  username: {
    label: t("register.byUsername.rules.username.label"),
    expression: /^[a-zA-Z0-9]{6,12}$/,
    message: t("register.byUsername.rules.username.message.pattern"),
  },
  password: {
    label: t("register.byUsername.rules.password.label"),
    expression: /^[a-zA-Z0-9]{6,12}$/,
    message: t("register.byUsername.rules.password.message.pattern"),
  },
  nickname: {
    label: t("register.byUsername.rules.nickname.label"),
    expression: /^[^@# ,.?，。？]{2,15}$/,
    message: t("register.byUsername.rules.nickname.message.pattern"),
  },
};

// 登录按钮是否禁用
const registerButtonDisabled = computed(() => {
  return !(
    regulation.username.expression.test(formValue.value.username) &&
    regulation.nickname.expression.test(formValue.value.nickname) &&
    regulation.password.expression.test(formValue.value.password)
  );
});

// 表单验证规则
const rules = {
  username: [
    {
      required: true,
      message: t("register.byUsername.rules.username.message.required"),
      trigger: "blur",
    },
    {
      pattern: regulation.username.expression,
      message: regulation.username.message,
      trigger: "blur",
    },
  ],
  nickname: [
    {
      required: true,
      message: t("register.byUsername.rules.nickname.message.required"),
      trigger: "blur",
    },
    {
      pattern: regulation.nickname.expression,
      message: regulation.nickname.message,
      trigger: "blur",
    },
  ],
  password: [
    {
      required: true,
      message: t("register.byUsername.rules.password.message.required"),
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
 * 注册
 */
const registerHandler = async () => {
  // 表单验证
  if (
    !(
      regulation.username.expression.test(formValue.value.username) &&
      regulation.nickname.expression.test(formValue.value.nickname) &&
      regulation.password.expression.test(formValue.value.password)
    )
  ) {
    return Message.error(t("register.byUsername.rules.rulesError"));
  }

  // 注册
  const {
    data: { code, msg },
  } = await userRequest.register({
    username: formValue.value.username,
    nickname: formValue.value.nickname,
    password: md5(formValue.value.password + salt), // 密码加密
  });

  if (code === 200) {
    Message.success(t("register.message.success")); // 提示注册成功
    return setTimeout(() => router.replace("/login"), 1000); // 1秒后返回上一页
  }

  Message.error(msg); // 提示注册失败
};
</script>

<template>
  <div class="flex justify-center items-center p-4">
    <n-el
      class="px-5 py-6 w-full rd-3 shadow translate-y-1/8"
      style="background: var(--register-panel-background-color)"
    >
      <title-bar class="mb-4" :title="t('register.title')" />
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
            :placeholder="t('register.byUsername.placeholder.username')"
            size="large"
          />
        </n-form-item>
        <n-form-item path="nickname" style="--n-label-height: 10px">
          <n-input
            v-model:value="formValue.nickname"
            style="
              --n-height: 50px;
              --n-border-radius: 0.75rem;
              --n-border-hover: var(--primary-color);
              --n-border-focus: var(--primary-color);
            "
            :placeholder="t('register.byUsername.placeholder.nickname')"
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
            :placeholder="t('register.byUsername.placeholder.password')"
            size="large"
          />
        </n-form-item>
      </n-form>

      <!-- 按钮 -->
      <div class="">
        <n-button
          type="primary"
          class="w-full h-12 rounded-xl text-5 shadow"
          :disabled="registerButtonDisabled"
          @click="registerHandler"
          >{{ t("register.byUsername.button.register") }}</n-button
        >
      </div>

      <div class="flex justify-center items-center mt-4 text-gray gap-2">
        <div>{{ t("register.byUsername.other.loginByPhone") }}</div>
        <div>|</div>
        <div @click="router.replace('login')">
          {{ t("register.byUsername.other.login") }}
        </div>
        <div>|</div>
        <div>{{ t("register.byUsername.other.forgotPassword") }}</div>
      </div>
    </n-el>
  </div>
</template>
