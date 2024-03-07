<script setup lang="ts">
import DefoGameMapMini from "@/components/defo-game-map-mini.vue";
import { defineProps, onMounted, ref } from "vue";
import { defalutAvatar } from "@/config";
import { userRequest } from "@/api/methods/user";
import { UserResp } from "@/types/user";

const props = defineProps<{
  detailData: any;
}>();

console.log(props.detailData);

const gameMode: any = {
  1: "练习",
  2: "排位",
  3: "对战",
};

const recordStatus: any = {
  1: "正常",
  2: "冻结",
  3: "删除",
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
              {{ gameMode[props.detailData.type] }}
            </n-descriptions-item>
            <n-descriptions-item label="阶数">
              {{ props.detailData.dimension }} 阶
            </n-descriptions-item>
            <n-descriptions-item label="耗时">
              {{ props.detailData.duration }}
            </n-descriptions-item>
            <n-descriptions-item label="步数">
              {{ props.detailData.step }}
            </n-descriptions-item>
            <n-descriptions-item label="记录状态">
              {{ recordStatus[props.detailData.status] }}
            </n-descriptions-item>
            <n-descriptions-item label="记录时间">
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
          <n-descriptions-item label="游戏模式">
            <div>{{ gameMode[props.detailData.type] }}</div>
          </n-descriptions-item>
          <n-descriptions-item label="阶数">
            {{ props.detailData.dimension }} 阶
          </n-descriptions-item>
          <n-descriptions-item label="耗时">
            {{ props.detailData.duration }}
          </n-descriptions-item>
          <n-descriptions-item label="步数">
            {{ props.detailData.step }}
          </n-descriptions-item>
          <n-descriptions-item label="记录时间" :span="2">
            {{ new Date(props.detailData.createdAt).toLocaleString() }}
          </n-descriptions-item>
        </n-descriptions>
        <n-descriptions bordered class="w-full">
          <n-descriptions-item label="打乱与还原">
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
:deep .n-descriptions.n-descriptions--bordered .n-descriptions-table-wrapper {
  border-radius: 0.5rem;
}
</style>
