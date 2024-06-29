<script setup lang="ts">
import TopBar from "@/components/top-bar.vue";
import { useUpdateLogStore } from "@/store/update-log";
import { onMounted } from "vue";

const updateLogStore = useUpdateLogStore();

onMounted(() => {
  updateLogStore.fetchLogs(
    {
      pagination: {
        page: 1,
        pageSize: 10,
      },
    },
    "replace"
  );
});
</script>

<template>
  <div>
    <top-bar title="更新日志" class="mt-4" />

    <div class="p-4">
      <!-- <div class="mt-2">内部测试期间不实行版本号管理</div> -->

      <div class="mt-4">
        <n-el
          class="p-4 mt-4 rounded-xl shadow-sm"
          style="background: var(--update-card-background-color)"
          v-for="(data, index) in updateLogStore.getLogs"
          :key="index"
        >
          <div class="flex">
            <div class="text-5 font-bold">v{{ data.version }}</div>
            <n-tag
              class="ml-4"
              :bordered="false"
              type="success"
              v-if="data.type === 1"
            >
              日常更新
            </n-tag>
            <n-tag
              class="ml-4"
              :bordered="false"
              type="error"
              v-if="data.type === 2"
            >
              重大更新
            </n-tag>
          </div>

          <div class="mt-2">
            <!-- 新增 -->
            <div v-if="data.content.new.length > 0">
              <div class="text-4">🚀 新特性</div>
              <ul class="list-disc">
                <li
                  class="ml-8"
                  v-for="(item, index) in data.content.new"
                  :key="index"
                >
                  {{ item }}
                </li>
              </ul>
            </div>

            <!-- 修复 -->
            <div v-if="data.content.fix.length > 0">
              <div class="text-4 mt-2">🐛 修复</div>
              <ul class="list-disc">
                <li
                  class="ml-8"
                  v-for="(item, index) in data.content.fix"
                  :key="index"
                >
                  {{ item }}
                </li>
              </ul>
            </div>

            <!-- 优化 -->
            <div v-if="data.content.optimize.length > 0">
              <div class="text-4 mt-2">🔧 优化</div>
              <ul class="list-disc">
                <li
                  class="ml-8"
                  v-for="(item, index) in data.content.optimize"
                  :key="index"
                >
                  {{ item }}
                </li>
              </ul>
            </div>
          </div>
        </n-el>
      </div>
    </div>
  </div>
</template>
