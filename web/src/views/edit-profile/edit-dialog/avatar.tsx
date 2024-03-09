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
import { defineComponent, ref } from "vue";
import { useI18n } from "vue-i18n";

export default defineComponent({
  name: "Avatar",
  props: {
    avatar: {
      type: String,
      default: "",
    },
    update: {
      type: Function,
      required: true,
    },
  },
  setup(props) {
    props.update?.({ name: "hello,ec" });

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
    return () => (
      <div class="flex flex-col items-center gap-4">
        <NAvatar
          class="rounded-full"
          size={128}
          object-fit="cover"
          src={currentAvatar.value || defaultAvatar}
          fallback-src={defaultAvatar}
        />
        <div class="flex justify-around items-center w-full">
          <NUpload
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
            <NButton>{t("editProfile.dialog.uploadLabel")}</NButton>
          </NUpload>
          <NButton
            type="primary"
            onClick={() => uploadRef.value?.submit()}
            disabled={!fileList.value.length}
          >
            {t("editProfile.dialog.avatar.submit")}
          </NButton>
        </div>
      </div>
    );
  },
});
