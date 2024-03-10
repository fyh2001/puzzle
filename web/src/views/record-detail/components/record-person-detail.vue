<script setup lang="ts">
import DefoGameMapMini from "@/components/defo-game-map-mini.vue";
import { defineProps, onMounted, ref } from "vue";
import { defaultAvatar } from "@/config";
import { userRequest } from "@/api/methods/user";
import { UserResp } from "@/types/user";
import { useI18n } from "vue-i18n";

const { t } = useI18n();

const props = defineProps<{
  detailData: any;
}>();

const gameMode: any = {
  1: t("recordDetail.content.gameMode.content.practice"),
  2: t("recordDetail.content.gameMode.content.rank"),
  3: t("recordDetail.content.gameMode.content.battle"),
};

const recordStatus: any = {
  1: t("recordDetail.content.status.content.normal"),
  2: t("recordDetail.content.status.content.freeze"),
  3: t("recordDetail.content.status.content.delete"),
};

const userInfo = ref<UserResp>();
const getUseInfo = async () => {
  const {
    data: { code, data: userInfoData },
  } = await userRequest.list({
    pagination: {
      page: 1,
      pageSize: 1,
    },
    id: props.detailData.userId,
  });

  if (code === 200) {
    userInfo.value = userInfoData.records[0];
  }
};

onMounted(() => {
  getUseInfo();
});
</script>

<template>
  <n-el class="h-full">
    <n-carousel :show-dots="false" :loop="false" :space-between="20">
      <!--   用户信息 和 记录信息-->
      <n-el class="flex flex-col items-center gap-4 w-full">
        <!-- 用户信息 -->
        <n-el
          class="flex flex-col items-center gap-4 p-6 w-full bg-white rounded-xl shadow"
          style="background-color: var(--user-card-background-color)"
        >
          <n-avatar
            class="shadow-lg"
            round
            object-fit="cover"
            :size="120"
            :src="userInfo?.avatar || defaultAvatar"
            :fallback-src="defaultAvatar"
          />
          <n-el class="text-gray-5 text-4 font-bold">{{
            userInfo?.nickname
          }}</n-el>
        </n-el>
        <!-- 记录信息 -->
        <n-el
          class="flex flex-col items-center gap-4 p-4 w-full bg-white rounded-xl shadow"
          style="background-color: var(--user-card-background-color)"
        >
          <n-descriptions
            :column="1"
            class="w-full"
            label-align="center"
            label-placement="left"
            bordered
          >
            <n-descriptions-item
              :label="t('recordDetail.content.gameMode.label')"
            >
              {{ gameMode[props.detailData.type] }}
            </n-descriptions-item>
            <n-descriptions-item
              :label="t('recordDetail.content.dimension.label')"
            >
              {{
                t("recordDetail.content.dimension.content", {
                  dimension: props.detailData.dimension,
                })
              }}
            </n-descriptions-item>
            <n-descriptions-item :label="t('recordDetail.content.duration')">
              {{ props.detailData.duration }}
            </n-descriptions-item>
            <n-descriptions-item :label="t('recordDetail.content.step')">
              {{ props.detailData.step }}
            </n-descriptions-item>
            <n-descriptions-item
              :label="t('recordDetail.content.status.label')"
            >
              {{ recordStatus[props.detailData.status] }}
            </n-descriptions-item>
            <n-descriptions-item :label="t('recordDetail.content.dateTime')">
              {{ new Date(props.detailData.createdAt).toLocaleString() }}
            </n-descriptions-item>
          </n-descriptions>
        </n-el>
      </n-el>
      <!-- 打乱与还原 -->
      <n-el class="flex flex-col items-center gap-2 w-full">
        <n-descriptions
          :column="2"
          class="w-full"
          label-align="center"
          label-placement="left"
          bordered
        >
          <n-descriptions-item
            :label="t('recordDetail.content.gameMode.label')"
          >
            <div>{{ gameMode[props.detailData.type] }}</div>
          </n-descriptions-item>
          <n-descriptions-item
            :label="t('recordDetail.content.dimension.label')"
          >
            {{
              t("recordDetail.content.dimension.content", {
                dimension: props.detailData.dimension,
              })
            }}
          </n-descriptions-item>
          <n-descriptions-item :label="t('recordDetail.content.duration')">
            {{ props.detailData.duration }}
          </n-descriptions-item>
          <n-descriptions-item :label="t('recordDetail.content.step')">
            {{ props.detailData.step }}
          </n-descriptions-item>
          <n-descriptions-item
            :label="t('recordDetail.content.dateTime')"
            :span="2"
          >
            {{ new Date(props.detailData.createdAt).toLocaleString() }}
          </n-descriptions-item>
        </n-descriptions>
        <n-descriptions bordered class="w-full">
          <n-descriptions-item
            :label="t('recordDetail.content.scrambleAndSolution')"
          >
            <DefoGameMapMini
              :dimension="props.detailData?.dimension"
              :scramble="props.detailData?.scramble"
              :solution="props.detailData?.solution"
              button-size="small"
              button-class="h-10 text-3.5"
            />
          </n-descriptions-item>
        </n-descriptions>
      </n-el>
    </n-carousel>
  </n-el>
</template>

<style scoped>
:deep(.n-descriptions.n-descriptions--bordered .n-descriptions-table-wrapper) {
  border-radius: 0.5rem;
}
</style>
