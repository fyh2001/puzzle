<script setup lang="tsx">
import ProfileItem from "@/views/edit-profile/components/profile-item.vue";
import { ref, computed } from "vue";
import { useI18n } from "vue-i18n";
import { defaultAvatar } from "@/config";
import { useUserStore } from "@/store/user";
import AvatarDialog from "./edit-dialog/avatar";
import NicknameDialog from "./edit-dialog/nickname.vue";

const { t } = useI18n();
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
    dialog: NicknameDialog
  },
  {
    label: t("editProfile.content.nickname"),
    content: userStore.getUser.nickname,
  },
]);

const commonUpdate = (form: Record<string, any>) => {
  console.log("common update: ", form);
};
</script>

<template>
  <div>
    <top-bar class="py-4" :title="t('editProfile.title')" />
    <div class="p-4 text-4.2">
      <div class="flex flex-col gap-6 px-6 py-4 bg-white rounded-xl">
        <profile-item
          v-for="(data, index) in profileItemOptions"
          :key="index"
          :label="data.label"
          :content="data.content"
          :avatar="data.avatar"
          @update="commonUpdate"
        >
          <template #dialog="dialogProps" v-if="data.dialog">
            <component :is="data.dialog" v-bind="dialogProps"/>
          </template>
        </ec-item>
      </div>
    </div>
  </div>
</template>

<style scoped></style>
