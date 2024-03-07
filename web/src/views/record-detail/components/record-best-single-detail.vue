<script setup lang="ts">
import DefoGameMapMini from "@/components/defo-game-map-mini.vue";
import { defineProps, onMounted, ref } from "vue";
import { defalutAvatar } from "@/config";
import { userRequest } from "api/user";
import { recordRequest } from "api/record";
import { UserResp } from "@/types/user";
import { RecordResp } from "@/types/record";
import { formatDurationInRecord } from "@/utils/time";

const props = defineProps<{
  detailData: any;
}>();

const gameMode: any = {
  1: "练习",
  2: "排位",
  3: "对战",
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

const recordDetail = ref<RecordResp>();
const getRecordDetail = async () => {
  const {
    data: { code, data: recordDetailData },
  } = await recordRequest.listRecord({
    pagination: {
      page: 1,
      pageSize: 1,
    },
    id: props.detailData.recordId,
  });

  if (code === 200) {
    recordDetail.value = recordDetailData.records[0];
    recordDetail.value.durationFormat = formatDurationInRecord(
      recordDetail.value.duration
    );
  }
};

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
            :size="120"
            :src="userInfo?.avatar || defalutAvatar"
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
            <n-descriptions-item label="游戏模式">
              {{ gameMode[recordDetail.type] }}
            </n-descriptions-item>
            <n-descriptions-item label="阶数">
              {{ props.detailData.dimension }} 阶
            </n-descriptions-item>
            <n-descriptions-item label="排名">
              第 {{ props.detailData.ranked }} 名
            </n-descriptions-item>
            <n-descriptions-item label="耗时">
              {{ props.detailData.duration }}
            </n-descriptions-item>
            <n-descriptions-item label="步数">
              {{ props.detailData.step }}
            </n-descriptions-item>
            <n-descriptions-item label="记录打破次数">
              {{ props.detailData.recordBreakCount }} 次
            </n-descriptions-item>
            <n-descriptions-item label="初次纪录">
              {{ new Date(props.detailData.createdAt).toLocaleString() }}
            </n-descriptions-item>
            <n-descriptions-item label="上次纪录">
              {{ new Date(props.detailData.updatedAt).toLocaleString() }}
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
          <n-descriptions-item label="游戏模式">
            <div>{{ gameMode[recordDetail.type] }}</div>
          </n-descriptions-item>
          <n-descriptions-item label="阶数">
            {{ recordDetail.dimension }} 阶
          </n-descriptions-item>
          <n-descriptions-item label="耗时">
            {{ recordDetail.durationFormat }}
          </n-descriptions-item>
          <n-descriptions-item label="步数">
            {{ recordDetail.step }}
          </n-descriptions-item>
          <n-descriptions-item label="记录时间" :span="2">
            {{ new Date(recordDetail.createdAt).toLocaleString() }}
          </n-descriptions-item>
        </n-descriptions>
        <n-descriptions bordered class="w-full">
          <n-descriptions-item label="打乱与还原">
            <DefoGameMapMini
              :dimension="recordDetail?.dimension"
              :scramble="recordDetail?.scramble"
              :solution="recordDetail?.solution"
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
:deep .n-descriptions.n-descriptions--bordered .n-descriptions-table-wrapper {
  border-radius: 0.5rem;
}
</style>
