import { onUnmounted, ref } from "vue";

export type Cleanup = () => any;
export type CallbackReturn = void | Cleanup;
export type Callback = (
  ...args: any[]
) => CallbackReturn | Promise<CallbackReturn>;

export const useIntervalAsync = (callback: Callback, delay: number) => {
  const timeout = ref<number | null>(null);
  const canceled = ref<boolean>(false);
  const cleanup = ref<Cleanup | void>();

  const run: TimerHandler = async () => {
    if (canceled.value) {
      return;
    }
    // 清理之前的回调函数
    if (typeof cleanup.value === "function") {
      cleanup.value();
    }
    // 执行回调函数并获取清理函数
    cleanup.value = await Promise.resolve(callback());
    // 设置下一次任务轮询的定时器
    timeout.value = globalThis.setTimeout(run, delay);
  };

  // 初始化任务轮询
  run();

  // 刷新任务轮询，取消当前定时器，重新执行回调函数
  const flush = () => {
    timeout.value && globalThis.clearTimeout(timeout.value);
    run();
  };

  // 取消任务轮询，清理定时器和回调函数
  const cancel = () => {
    timeout.value && globalThis.clearTimeout(timeout.value);
    canceled.value = true;
    if (typeof cleanup.value === "function") {
      cleanup.value();
    }
  };

  // 恢复任务轮询，重新启动定时器
  const recover = () => {
    canceled.value = false;
    flush();
  };

  // 在组件卸载时取消任务轮询
  onUnmounted(() => {
    cancel();
  });

  return {
    flush,
    cancel,
    recover,
  };
};

export default useIntervalAsync;
