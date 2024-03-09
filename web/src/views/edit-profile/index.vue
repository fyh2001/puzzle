<script setup lang="tsx">
import ProFileItem from "@/views/edit-profile/components/profile-item.vue";
import { ref, computed } from "vue";
import { useI18n } from "vue-i18n";
import { defaultAvatar } from "@/config";
import { useUserStore } from "@/store/user";
import { useDialog, useMessage } from "naive-ui";
import { baseURL } from "@/api";
import type { UploadFileInfo, UploadInst } from "naive-ui";

const { t } = useI18n();
const Dialog = useDialog();
const Message = useMessage();
const userStore = useUserStore();

const profileItemOptions = computed(() => [
  {
    label: t("editProfile.content.avatar"),
    avatar: userStore.getUser.avatar || defaultAvatar,
    onClick: handleAvatarDialogVisible,
  },
  {
    label: t("editProfile.content.username"),
    content: userStore.getUser.username,
  },
  {
    label: t("editProfile.content.nickname"),
    content: userStore.getUser.nickname,
  },
]);

const formData = ref({
  avatar: "",
});

const uploadRef = ref<UploadInst | null>(null);
const fileList = ref<UploadFileInfo[]>([]);
const handleAvatarDialogVisible = async () => {
  const header = { Authorization: userStore.getToken };

  const handleUpdateAvatar = (file: UploadFileInfo) => {
    formData.value.avatar = userStore.user.avatar;

    // 判断文件是否超过5Mb
    if (file.file && file.file.size > 5 * 1024 * 1024) {
      return Message.error("文件大小不能超过5Mb");
    }

    fileList.value = [file];

    const reader = new FileReader();
    reader.onload = () => (formData.value.avatar = reader.result as string);
    if (file.file) reader.readAsDataURL(file.file);
  };

  const handleUploadFinish = (file: UploadFileInfo) => {
    if (file.status === "finished") {
      userStore.fetchUser();
      Message.success(t("editProfile.messgae.avatar.success"));
      Dialog.destroyAll();
    }
  };

  const handleUploadError = (err: UploadFileInfo) => {
    console.error(err);
  };

  Dialog.warning({
    title: t("editProfile.dialog.avatar.label"),
    content: () => {
      return (
        <div class="flex flex-col items-center gap-4">
          <n-avatar
            class="rounded-full"
            size={128}
            object-fit="cover"
            src={formData.value.avatar || defaultAvatar}
            fallback-src={defaultAvatar}
          />
          <div class="flex justify-around items-center w-full">
            <n-upload
              ref={uploadRef}
              accept="image/*,.png,.jpg,.jpeg,.svg"
              name="avatar"
              headers={header}
              action={baseURL + "/user/update-avatar"}
              file-list={fileList.value}
              on-update:file-list={(files: UploadFileInfo[]) =>
                handleUpdateAvatar(files.pop() as UploadFileInfo)
              }
              on-finish={({ file }: { file: UploadFileInfo }) =>
                handleUploadFinish(file)
              }
              on-error={(err: UploadFileInfo) => handleUploadError(err)}
              default-upload={false}
              show-file-list={false}
              multiple={false}
            >
              <n-button>{t("editProfile.dialog.uploadLabel")}</n-button>
            </n-upload>
            <n-button
              type="primary"
              onClick={() => uploadRef.value?.submit()}
              disabled={!fileList.value.length}
            >
              {t("editProfile.dialog.avatar.submit")}
            </n-button>
          </div>
        </div>
      );
    },
  });
};
</script>

<template>
  <div>
    <top-bar class="py-4" :title="t('editProfile.title')" />
    <div class="p-4 text-4.2">
      <div class="flex flex-col gap-6 px-6 py-4 bg-white rounded-xl">
        <pro-file-item
          v-for="(data, index) in profileItemOptions"
          :key="index"
          :label="data.label"
          :content="data.content"
          :avatar="data.avatar"
          @click="data.onClick"
        />
      </div>
    </div>
  </div>
</template>

<style scoped></style>
