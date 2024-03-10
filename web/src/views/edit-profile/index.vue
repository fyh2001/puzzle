<script setup lang="tsx">
import ProfileItem from "@/views/edit-profile/components/profile-item.vue";
import { computed } from "vue";
import { useI18n } from "vue-i18n";
import { defaultAvatar } from "@/config";
import { useUserStore } from "@/store/user";
import AvatarDialog from "@/views/edit-profile/dialog/avatar.vue";
import NicknameDialog from "@/views/edit-profile/dialog/nickname.vue";
import { userRequest } from "@/api/methods/user";
import { useMessage, useDialog } from "naive-ui";
import type { UserModel } from "@/types/user";

const { t } = useI18n();
const Message = useMessage();
const Dialog = useDialog();
const userStore = useUserStore();

const profileItemOptions = computed(() => [
  {
    label: t("editProfile.content.avatar"),
    avatar: userStore.getUser.avatar || defaultAvatar,
    dialog: AvatarDialog,
  },
  {
    label: t("editProfile.content.username"),
    content: userStore.getUser.username,
  },
  {
    label: t("editProfile.content.nickname"),
    content: userStore.getUser.nickname,
    dialog: NicknameDialog,
  },
]);

// 公共的更新函数
const commonUpdate = async (form: UserModel) => {
  console.log("common update: ", form);

  form.id = userStore.getUser.id;

  const {
    data: { code, msg },
  } = await userRequest.update(form);

  if (code === 200) {
    userStore.fetchUser();
    Message.success(t("editProfile.messgae.success"));
  } else {
    Message.error(msg);
  }
  window.history.back();
  Dialog.destroyAll();
};
</script>

<template>
  <n-el>
    <top-bar class="py-4" :title="t('editProfile.title')" />
    <n-el class="p-4 text-4.2">
      <div
        class="flex flex-col gap-6 px-6 py-4 rounded-xl"
        style="background: var(--edit-card-background-color)"
      >
        <profile-item
          v-for="(data, index) in profileItemOptions"
          :key="index"
          :label="data.label"
          :content="data.content"
          :avatar="data.avatar"
          @update="commonUpdate"
        >
          <template #dialog v-if="data.dialog">
            <component :is="data.dialog" :update="commonUpdate" />
          </template>
        </profile-item>
      </div>
    </n-el>
  </n-el>
</template>

<style scoped></style>
