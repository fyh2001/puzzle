import { onMounted, onUnmounted, ref } from "vue";
import * as _ from "lodash";

export function useScrollTo(
  threshold: number,
  callback: Function,
  delay: number = 2000,
  target?: string | Element
) {
  const targetElement = ref<Element | null>(null);

  const handleScroll = _.throttle(() => {
    let distance: number;
    if (targetElement.value) {
      const rect = targetElement.value.getBoundingClientRect();
      distance = rect.top;
    } else {
      const scrollTop =
        document.documentElement.scrollTop || document.body.scrollTop;
      const windowHeight =
        document.documentElement.clientHeight || document.body.clientHeight;
      const scrollHeight =
        document.documentElement.scrollHeight || document.body.scrollHeight;
      distance = scrollHeight - scrollTop - windowHeight;
    }

    if (distance <= threshold) {
      callback();
    }
  }, delay);

  onMounted(() => {
    if (typeof target === "string") {
      targetElement.value = document.querySelector(target);
    } else if (target instanceof Element) {
      targetElement.value = target;
    }

    window.addEventListener("scroll", handleScroll);
  });

  onUnmounted(() => {
    window.removeEventListener("scroll", handleScroll);
  });
}
