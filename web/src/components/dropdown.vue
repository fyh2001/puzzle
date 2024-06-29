<script setup lang="tsx">
import {
  NotificationsActiveRound,
  WbSunnyRound,
  SettingsRound,
  EventNoteRound,
  DarkModeRound,
  AttachEmailOutlined,
  StorefrontRound,
  BookmarkRound,
  LanguageRound,
} from "@vicons/material";
import router from "@/routers/index";
import { useThemeStore } from "@/store/theme";
import { useCommonStore } from "@/store/common";
import { useNotificationStore } from "@/store/notification";
import { computed } from "vue";
import { useI18n } from "vue-i18n";

const { t, locale } = useI18n();

const themeStore = useThemeStore();
const commonStore = useCommonStore();
const notificationStore = useNotificationStore();

const emit = defineEmits(["select"]);

const props = defineProps<{
  content?: string;
  options: any[];
  showDivider: boolean;
}>();

const options = computed(() => [
  ...props.options,
  // 分割线
  {
    key: "divider",
    show: props.options.length > 0 && props.showDivider,
    type: "divider",
  },
  // 通知
  {
    label: t("dropdown.notification.label"),
    key: "notification",
    disabled: false,
    icon: () => {
      return (
        <n-badge
          dot
          processing
          value={notificationStore.getNotificationUnreadTotal}
          show={notificationStore.getNotificationUnreadTotal > 0}
          max={10}
        >
          <n-el class="flex items center" style="color: var(--primary-color)">
            <n-icon size="18" component={NotificationsActiveRound} />
          </n-el>
        </n-badge>
      );
    },
  },
  // 深色模式
  {
    label: t("dropdown.darkMode.label"),
    key: "dark",
    disabled: false,
    icon: () => {
      return (
        <n-el class="flex items-center" style="color: var(--primary-color)">
          <n-icon size="18" component={DarkModeRound} />
        </n-el>
      );
    },
    children: [
      {
        label: t("dropdown.darkMode.on"),
        key: "darkOn",
        disabled: false,
        icon: () => {
          return (
            <n-el class="flex items-center" style="color: var(--primary-color)">
              <n-icon
                size="18"
                class="text-blue-500"
                component={DarkModeRound}
              />
            </n-el>
          );
        },
      },
      {
        label: t("dropdown.darkMode.off"),
        key: "darkOff",
        disabled: false,
        icon: () => {
          return (
            <n-el class="flex items-center" style="color: var(--primary-color)">
              <n-icon
                size="18"
                class="text-yellow-500"
                component={WbSunnyRound}
              />
            </n-el>
          );
        },
      },
    ],
  },
  // 主题切换
  {
    label: t("dropdown.theme.label"),
    key: "themeColor",
    disabled: false,
    icon: () => {
      return (
        <n-el class="flex items-center" style="color: var(--primary-color)">
          <n-icon size="18" component={StorefrontRound} />
        </n-el>
      );
    },
    children: themeStore.loadedThemes.map((theme) => ({
      label: t(`dropdown.theme.color.${theme.name}`),
      key: theme.name,
      disabled: false,
      icon: () => (
        <n-icon
          size="18"
          class={theme.menuIconClass}
          component={BookmarkRound}
        />
      ),
    })),
  },
  // 语言切换
  {
    label: t("dropdown.language.label"),
    key: "language",
    icon: () => {
      return (
        <n-el class="flex items-center" style="color: var(--primary-color)">
          <n-icon size="18" component={LanguageRound} />
        </n-el>
      );
    },
    children: [
      {
        label: t("dropdown.language.content.cn"),
        key: "cn",
        disabled: false,
      },
      {
        label: t("dropdown.language.content.us"),
        key: "us",
        disabled: false,
      },
      {
        label: t("dropdown.language.content.jp"),
        key: "jp",
        disabled: false,
      },
    ],
  },
  // 设置
  {
    label: t("dropdown.setting.label"),
    key: "setting",
    disabled: true,
    icon: () => {
      return (
        <n-el class="flex items-center" style="color: var(--primary-color)">
          <n-icon size="18" component={SettingsRound} />
        </n-el>
      );
    },
  },
  // 更新日志
  {
    label: t("dropdown.updateLog.label"),
    key: "updateLog",
    disabled: false,
    icon: () => {
      return (
        <n-el class="flex items-center" style="color: var(--primary-color)">
          <n-icon size="18" component={EventNoteRound} />
        </n-el>
      );
    },
  },
  // 反馈
  {
    label: t("dropdown.feedback.label"),
    key: "feedback",
    disabled: true,
    icon: () => {
      return (
        <n-el class="flex items-center" style="color: var(--primary-color)">
          <n-icon size="18" component={AttachEmailOutlined} />
        </n-el>
      );
    },
  },
]);

const handleSelect = (key: string) => {
  const func = {
    notification: () => router.push("/notification"),
    darkOn: () => themeStore.openDarkMode(),
    darkOff: () => themeStore.closeDarkMode(),
    cn: () => {
      locale.value = "cn";
      commonStore.changeLanguage("cn");
    },
    us: () => {
      locale.value = "us";
      commonStore.changeLanguage("us");
    },
    jp: () => {
      locale.value = "jp";
      commonStore.changeLanguage("jp");
    },
    // 主题切换
    ...themeStore.loadedThemes.reduce((obj, theme) => {
      obj[theme.name] = () => themeStore.changeThemeColor(theme.name);
      return obj;
    }, {} as Record<string, () => any>),

    updateLog: () => router.push("/update-log"),
  }[key];
  func ? func() : emit("select", key);
};
</script>

<template>
  <div>
    <n-badge
      processing
      :value="notificationStore.getNotificationUnreadTotal"
      :max="9"
    >
      <n-dropdown
        placement="bottom-end"
        trigger="click"
        size="large"
        :options="options"
        @select="handleSelect"
      >
        <n-button strong secondary type="primary">
          {{ content }}
        </n-button>
      </n-dropdown>
    </n-badge>
  </div>
</template>
