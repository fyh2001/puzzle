<script lang="ts" setup>
import { computed, ref } from "vue";
import { NInput, NForm, NFormItem } from "naive-ui";
import { useI18n } from "vue-i18n";

const props = defineProps<{
  update: Function;
}>();

const { t } = useI18n();

const formData = ref({
  nickname: "",
});

const rules = {
  nickname: {
    pattern:
      /^(?!.*[@# ,.?，。？])([\u4e00-\u9fa5]{2,6}|[a-zA-Z]{2,10}|\d{11}|(?=.*[a-zA-Z0-9].*)([\u4e00-\u9fa5a-zA-Z0-9]{2,8}))$/,
    message: t("register.byUsername.rules.nickname.message.pattern"),
    trigger: ["blur", "input", "focus"],
  },
};

const buttonDisabled = computed(() => {
  return !rules.nickname.pattern.test(formData.value.nickname);
});
</script>

<template>
  <div class="mt-4">
    <NForm ref="formRef" :model="formData" :rules="rules">
      <NFormItem
        path="nickname"
        label-width="100"
        style="--n-label-height: 0px"
      >
        <NInput
          v-model:value="formData.nickname"
          :placeholder="t('editProfile.dialog.nickname.input.placeholder')"
          style="
            --n-border-radius: 0.5rem;
            --n-border-hover: var(--primary-color);
            --n-border-focus: var(--primary-color);
          "
        />
      </NFormItem>
      <div class="flex justify-end items-center gap-4 mt-4">
        <n-button
          strong
          secondary
          type="primary"
          class="rounded-xl text-4 shadow"
          :disabled="buttonDisabled"
          @click="props.update(formData)"
        >
          {{ t("editProfile.dialog.nickname.button.confirm") }}
        </n-button>
      </div>
    </NForm>
  </div>
</template>

<style scoped>
:deep(.n-form-item .n-form-item-feedback-wrapper) {
  min-height: 0;
}
</style>
