<script lang="ts" setup>
import TopBar from "@/components/top-bar.vue";
import ByUsername from "@/views/login/components/by-username.vue";
import { useI18n } from "vue-i18n";

const { t } = useI18n();
</script>

<template>
  <div>
    <top-bar class="my-4 ml-2" :title="t('login.title')" />
    <by-username />
  </div>
</template>
