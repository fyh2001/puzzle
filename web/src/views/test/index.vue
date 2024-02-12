<template>
  <div>
    <n-qr-code :value="text" />
  </div>
</template>

<script lang="ts" setup>
const text =
  "otpauth://totp/puzzle:defo1215.cn?algorithm=SHA1&digits=6&issuer=puzzle&period=30&secret=UOIA6VWYKYTUDJEGUJNLLD55L5E3UTWC";
</script>
