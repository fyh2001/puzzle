<script setup lang="ts">
import DefoGameMapMini from "@/components/defo-game-map-mini.vue";
import { defineProps, onMounted, ref } from "vue";
import { defaultAvatar } from "@/config";
import { userRequest } from "api/user";
import { recordRequest } from "api/record";
import { UserResp } from "@/types/user";
import { RecordResp } from "@/types/record";
import { formatDurationInRecord } from "@/utils/time";
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

const recordDetail = ref<RecordResp[]>();
const getRecordDetail = async () => {
  const {
    data: { code, data: recordDetailData },
  } = await recordRequest.listRecord({
    pagination: {
      page: 1,
      pageSize: props.detailData.type,
    },
    ids: props.detailData.recordIds.split(","),
  });

  if (code === 200) {
    recordDetail.value = recordDetailData.records.map((item: any) => {
      item.durationFormat = formatDurationInRecord(item.duration);
      return item;
    });
  }
};

// const carouselDotVisible = ref(false);
// let carouselTimer: any = null;
// const handleCarouselChange = () => {
//   clearTimeout(carouselTimer);
//   carouselDotVisible.value = true;

//   carouselTimer = setTimeout(() => {
//     carouselDotVisible.value = false;
//   }, 1000);
// };

onMounted(() => {
  getUseInfo();
  getRecordDetail();
});
</script>

<template>
  <n-el class="h-full" v-if="recordDetail">
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
            <n-descriptions-item :label="t('recordDetail.content.type.label')">
              {{
                t("recordDetail.content.type.content.bestAverage", {
                  type: props.detailData.type,
                })
              }}
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
            <n-descriptions-item :label="t('recordDetail.content.rank.label')">
              {{
                t("recordDetail.content.rank.content", {
                  rank: props.detailData.ranked,
                })
              }}
            </n-descriptions-item>
            <n-descriptions-item :label="t('recordDetail.content.duration')">
              {{ props.detailData.duration }}
            </n-descriptions-item>
            <n-descriptions-item
              :label="t('recordDetail.content.recordBreakCount.label')"
            >
              {{
                t(
                  "recordDetail.content.recordBreakCount.content",
                  {
                    recordBreakCount: props.detailData.recordBreakCount,
                  },
                  props.detailData.recordBreakCount
                )
              }}
            </n-descriptions-item>
            <n-descriptions-item :label="t('recordDetail.content.createdAt')">
              {{ new Date(props.detailData.createdAt).toLocaleString() }}
            </n-descriptions-item>
            <n-descriptions-item :label="t('recordDetail.content.updatedAt')">
              {{ new Date(props.detailData.updatedAt).toLocaleString() }}
            </n-descriptions-item>
          </n-descriptions>
        </n-el>
      </n-el>
      <!-- 打乱与还原 -->
      <n-el
        class="flex flex-col items-center gap-2 w-full"
        v-for="(data, index) in recordDetail"
      >
        <n-descriptions
          :column="2"
          :title="
            t('recordDetail.content.detailDataTitle', { index: index + 1 })
          "
          class="w-full"
          label-align="center"
          label-placement="left"
          bordered
        >
          <n-descriptions-item
            :label="t('recordDetail.content.gameMode.label')"
          >
            <div>{{ gameMode[data.type] }}</div>
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
            {{ data.durationFormat }}
          </n-descriptions-item>
          <n-descriptions-item :label="t('recordDetail.content.step')">
            {{ data.step }}
          </n-descriptions-item>
          <n-descriptions-item
            :label="t('recordDetail.content.dateTime')"
            :span="2"
          >
            {{ new Date(data.createdAt).toLocaleString() }}
          </n-descriptions-item>
        </n-descriptions>
        <n-descriptions bordered class="w-full">
          <n-descriptions-item
            :label="t('recordDetail.content.scrambleAndSolution')"
          >
            <DefoGameMapMini
              :dimension="data?.dimension"
              :scramble="data?.scramble"
              :solution="data?.solution"
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

.v-enter-active,
.v-leave-active {
  transition: opacity 0.5s;
}

.v-enter-from,
.v-leave-to {
  opacity: 0;
}
</style>
