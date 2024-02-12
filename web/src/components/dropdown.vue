<script setup lang="tsx">
import {
  WbSunnyRound,
  SettingsRound,
  EventNoteRound,
  DarkModeRound,
  AttachEmailOutlined,
  StorefrontRound,
  BookmarkRound,
} from "@vicons/material";
import router from "@/routers/index";
import { useThemeStore } from "@/store/theme";
import { computed } from "vue";

const themeStore = useThemeStore();

const emit = defineEmits(["select"]);

const props = defineProps<{
  content?: string;
  options: any[];
  showDivider: boolean;
}>();

const options = computed(() => [
  ...props.options,
  {
    key: "divider",
    show: props.options.length > 0 && props.showDivider,
    type: "divider",
  },
  {
    label: "深色模式",
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
        label: "开启",
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
        label: "关闭",
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
  {
    label: "主题色",
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
      label: theme.label,
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
  {
    label: "设置",
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
  {
    label: "更新日志",
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
  {
    label: "用户反馈",
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
    darkOn: () => themeStore.openDarkMode(),
    darkOff: () => themeStore.closeDarkMode(),
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
  </div>
</template>
