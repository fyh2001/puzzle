<script lang="ts" setup>
import ByUsername from "@/views/register/components/by-username.vue";
import TopBar from "@/components/top-bar.vue";
import { useI18n } from "vue-i18n";

const { t } = useI18n();
</script>

<template>
  <div>
    <top-bar class="my-4 ml-2" :title="t('register.title')" />
    <by-username />
  </div>
</template>
