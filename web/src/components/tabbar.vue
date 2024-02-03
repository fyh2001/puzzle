<script lang="ts" setup>
import { computed, nextTick, onMounted, ref, watch } from "vue";
import { useRoute, useRouter } from "vue-router";

export interface TabbarItem {
  label: string;
  icon: any;
  path: string;
}

export interface TabItemOffset {
  left: number;
  width: number;
}

const route = useRoute();
const router = useRouter();

const props = defineProps<{
  tabsContent: TabbarItem[];
}>();

// tabbar的dom节点
const tabItemRef = ref();

// 盒子位置
const boxRect = ref<TabItemOffset>({
  width: 0,
  left: 0,
});

// 当前下标
const tabIndex = computed({
  get: () => props.tabsContent.findIndex((item) => item.path === route.path),
  set: (value) => router.replace(props.tabsContent[value].path),
});

onMounted(() =>
  // 监听tabIndex的变化，更新盒子位置
  watch(
    tabIndex,
    () =>
      nextTick(() => {
        const tabItem: HTMLDivElement =
          tabItemRef.value!.children[tabIndex.value]; // 获取选中的tab

        boxRect.value = {
          width: tabItem!.offsetWidth, // 获取选中的tab的宽度
          left: tabItem!.offsetLeft, // 获取选中的tab的偏移量
        };
      }),
    { immediate: true }
  )
);
</script>

<template>
  <n-el
    class="w-full py-3 bg-[#99CC33] bg-indigo-3 rounded-4 overflow-hidden"
    style="background: var(--tabbar-background-color)"
  >
    <!-- 选择框 -->
    <n-el
      class="absolute left-0 h-10 rounded-xl bg-[#669933] bg-indigo-5 transition-all duration-300"
      style="background: var(--tabar-select-box-color)"
      :style="{
        width: boxRect.width + 'px',
        transform: 'translateX(' + boxRect.left + 'px)',
      }"
    />
    <!-- 图标与标题 -->
    <div ref="tabItemRef" class="flex justify-around items-center">
      <n-el
        class="flex items-center px-3 py-1 gap-2 text-white font-bold z-100"
        v-for="(data, index) in tabsContent"
        :key="index"
        @click="tabIndex = index"
      >
        <n-icon class="text-3xl" :component="data.icon" />
        <n-el
          class="text-sm transition-all duration-300"
          :class="
            tabIndex === index ? 'opacity-100' : 'opacity-0 -translate-x-1/2'
          "
        >
          {{ tabIndex === index ? data.label : "" }}
        </n-el>
      </n-el>
    </div>
  </n-el>
</template>
