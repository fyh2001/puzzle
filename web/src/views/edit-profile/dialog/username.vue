<script lang="ts" setup>
import { computed, ref } from "vue";
import { NInput, NForm, NFormItem } from "naive-ui";
import { useI18n } from "vue-i18n";

const props = defineProps<{
  update: Function;
}>();

const { t } = useI18n();

const formData = ref({
  username: "",
});

const rules = {
  username: {
    pattern: /^[a-zA-Z0-9]{6,12}$/,
    message: t("register.byUsername.rules.username.message.pattern"),
    trigger: ["blur", "input", "focus"],
  },
};

const buttonDisabled = computed(() => {
  return !rules.username.pattern.test(formData.value.username);
});
</script>

<template>
  <div class="mt-4">
    <NForm ref="formRef" :model="formData" :rules="rules">
      <NFormItem
        path="username"
        label-width="100"
        style="--n-label-height: 0px"
      >
        <NInput
          v-model:value="formData.username"
          :placeholder="t('editProfile.dialog.username.input.placeholder')"
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
          {{ t("editProfile.dialog.username.button.confirm") }}
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
