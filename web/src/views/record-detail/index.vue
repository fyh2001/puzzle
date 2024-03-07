<script setup lang="ts">
import TopBar from "@/components/top-bar.vue";
import RecordPersonDetail from "@/views/record-detail/components/record-person-detail.vue";
import RecordBestSingleDetail from "@/views/record-detail/components/record-best-single-detail.vue";
import RecordBestAverageDetail from "@/views/record-detail/components/record-best-average-detail.vue";
import RecordBestStepDetail from "@/views/record-detail/components/record-best-step-detail.vue";
import { onMounted, ref } from "vue";
import { useRoute } from "vue-router";

const route = useRoute();

// 详情数据
const data = ref(null);
// 类型
const type = ref<string>("");

const init = () => {
  const { data: rowData, type: detailType } = route.query;
  data.value = JSON.parse(rowData as string);
  type.value = detailType as string;
};

onMounted(() => {
  init();
});
</script>

<template>
  <div class="flex flex-col h-screen">
    <top-bar class="pt-2" title="记录详情" />

    <div class="flex-grow p-4" v-if="data">
      <record-person-detail :detailData="data" v-if="type === 'person'" />
      <record-best-single-detail
        :detailData="data"
        v-else-if="type === 'bestSingle'"
      />
      <record-best-average-detail
        :detailData="data"
        v-else-if="type === 'bestAverage'"
      />
      <record-best-step-detail
        :detailData="data"
        v-else-if="type === 'bestStep'"
      />
    </div>
  </div>
</template>
