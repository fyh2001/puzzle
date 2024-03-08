<script lang="tsx" setup>
import {
  HealthAndSafetyRound,
  TransitEnterexitRound,
  AccountCircleRound,
} from "@vicons/material";
import Dropdown from "@/components/dropdown.vue";
import { useUserStore } from "@/store/user";
import router from "@/routers/index";
import { useDialog, useMessage } from "naive-ui";
import { defalutAvatar } from "@/config/index";
import { useI18n } from "vue-i18n";
import { computed } from "vue";

const { t } = useI18n();

const Dialog = useDialog();
const Message = useMessage();

const userStore = useUserStore();

const options = computed(() => [
  {
    label: t("mine.dropdown.content.editProfile"),
    key: "editProfile",
    disabled: false,
    show: userStore.getToken !== null,
    icon: () => {
      return (
        <n-el class="flex items-center" style="color: var(--primary-color)">
          <n-icon size="18" component={AccountCircleRound} />
        </n-el>
      );
    },
    props: {
      onClick: () => {
        router.push("/user-edit-profile");
      },
    },
  },
  {
    label: t("mine.dropdown.content.editPassword"),
    key: "changePassword",
    disabled: true,
    show: userStore.getToken !== null,
    icon: () => {
      return (
        <n-el class="flex items-center" style="color: var(--primary-color)">
          <n-icon size="18" component={HealthAndSafetyRound} />
        </n-el>
      );
    },
  },
  {
    label: t("mine.dropdown.content.logout.label"),
    key: "logout",
    // disabled: true,
    show: userStore.getToken !== null,
    icon: () => {
      return (
        <n-el class="flex items-center" style="color: var(--primary-color)">
          <n-icon size="18" component={TransitEnterexitRound} />
        </n-el>
      );
    },
    props: {
      onClick: () => {
        Dialog.warning({
          title: t("mine.dropdown.content.logout.title"),
          content: t("mine.dropdown.content.logout.content"),
          positiveText: t("mine.dropdown.content.logout.confirm"),
          negativeText: t("mine.dropdown.content.logout.cancel"),
          onPositiveClick: () => {
            userStore.logout();
            Message.success(t("mine.dropdown.content.logout.message.success"));

            setTimeout(() => {
              window.location.reload();
            }, 1000);
          },
        });
      },
    },
  },
]);
</script>

<template>
  <div class="p-4 w-screen">
    <div class="flex justify-between items-center">
      <title-bar :title="t('mine.title')" />

      <dropdown
        :content="t('mine.dropdown.label')"
        :options="options"
        :show-divider="userStore.getToken !== null"
      />
    </div>

    <n-el
      class="relative py-10 mt-4 w-full rounded-2xl shadow bg-cover overflow-hidden"
      style="background: var(--user-card-background-color)"
    >
      <!-- 已登录 -->
      <div
        class="relative flex justify-around items-center z-100"
        v-if="userStore.getToken"
      >
        <!-- 头像与昵称 -->
        <transition name="left" appear>
          <div class="text-center">
            <n-avatar
              class="mb-6 shadow"
              object-fit="cover"
              round
              :size="100"
              :src="
                userStore.getUser.avatar !== ''
                  ? userStore.getUser.avatar
                  : defalutAvatar
              "
              :fallback-src="defalutAvatar"
            />
            <div class="text-5">{{ userStore.getUser.nickname }}</div>
          </div>
        </transition>

        <!-- 最佳成绩与排名
        <transition name="right" appear>
          <div
            class="grid grid-cols-1 gap-2 text-4"
            v-if="userStore.userBestRecord"
          >
            <div>
              最佳单次:
              {{
                userStore.userBestRecord.bestSingle == ""
                  ? "暂无"
                  : userStore.getBestSingleRecord
              }}
            </div>
            <div>
              最佳平均:
              {{
                userStore.userBestRecord.bestAverage == ""
                  ? "暂无"
                  : userStore.getBestAverageRecord
              }}
            </div>
            <div>
              最佳排名:
              {{
                userStore.userBestRecord.bestRank == 0
                  ? "暂未开启"
                  : userStore.userBestRecord.bestRank
              }}
            </div>
          </div>
        </transition> -->
      </div>

      <!-- 未登录 -->
      <div class="flex justify-around items-center z-100" v-else>
        <!-- 头像与昵称 -->
        <transition name="left" appear>
          <div class="text-center" @click="router.push('login')">
            <n-avatar
              class="mb-6 shadow"
              :size="100"
              round
              :src="defalutAvatar"
              :fallback-src="defalutAvatar"
            />
            <div class="text-5">{{ t("mine.login") }}</div>
          </div>
        </transition>
      </div>
    </n-el>
  </div>
</template>

<style scoped>
.left-enter-active,
.left-leave-active {
  transition: all 1s;
}

.left-enter-from {
  opacity: 0;
  transform: translateX(50%);
}

.right-enter-active,
.right-leave-active {
  transition: all 1s;
}

.right-enter-from {
  opacity: 0;
  transform: translateX(-50%);
}
</style>
