<script lang="ts" setup>
import { useUserStore } from "@/store/user";
import { baseURL } from "@/api";
import { defaultAvatar } from "@/config";
import {
  NButton,
  NAvatar,
  NUpload,
  UploadFileInfo,
  useDialog,
  useMessage,
} from "naive-ui";
import { ref } from "vue";
import { useI18n } from "vue-i18n";

const props = withDefaults(
  defineProps<{
    avatar?: string;
    update: Function;
  }>(),
  {
    avatar: "",
  }
);

const userStore = useUserStore();

const { t } = useI18n();
const Message = useMessage();
const Dialog = useDialog();

const uploadRef = ref<any>(null);
const currentAvatar = ref(props.avatar || userStore.user.avatar);
const fileList = ref([] as UploadFileInfo[]);

const header = { Authorization: userStore.getToken || "" };

const handleUpdateAvatar = (file: UploadFileInfo) => {
  currentAvatar.value = userStore.user.avatar;

  // 判断文件是否超过5Mb
  if (file.file && file.file.size > 5 * 1024 * 1024) {
    return Message.error("文件大小不能超过5Mb");
  }

  // 更新文件列表
  fileList.value = [file];

  // 读取文件并显示
  const reader = new FileReader();
  reader.onload = () => (currentAvatar.value = reader.result as string);
  if (file.file) reader.readAsDataURL(file.file);
};

const handleUploadFinish = ({
  file,
  event,
}: {
  file: UploadFileInfo;
  event?: ProgressEvent;
}) => {
  console.log(event);

  if (file.status === "finished") {
    userStore.fetchUser();
    Message.success(t("editProfile.messgae.success"));
  } else {
    Message.error(t("editProfile.messgae.error"));
  }

  Dialog.destroyAll();
};

const handleUploadError = ({
  file,
  event,
}: {
  file: UploadFileInfo;
  event?: ProgressEvent;
}) => {
  console.error(file, event);
};
</script>

<template>
  <div class="flex flex-col items-center gap-4">
    <NAvatar
      class="rounded-full"
      :size="128"
      object-fit="cover"
      :src="currentAvatar || defaultAvatar"
      :fallback-src="defaultAvatar"
    />
    <div class="flex justify-around items-center w-full">
      <NUpload
        ref="uploadRef"
        accept="image/*,.png,.jpg,.jpeg,.svg"
        name="avatar"
        :headers="header"
        :action="baseURL + '/user/update-avatar'"
        :file-list="fileList"
        :default-upload="false"
        :show-file-list="false"
        :multiple="false"
        @update:file-list="(files: UploadFileInfo[]) => handleUpdateAvatar(files.pop() as UploadFileInfo)"
        @finish="handleUploadFinish"
        @error="handleUploadError"
      >
        <NButton>{{ t("editProfile.dialog.uploadLabel") }}</NButton>
      </NUpload>
      <NButton
        strong
        secondary
        type="primary"
        @click="() => uploadRef?.submit()"
        :disabled="!fileList.length"
      >
        {{ t("editProfile.dialog.avatar.button.confirm") }}
      </NButton>
    </div>
  </div>
</template>

<style scoped></style>
