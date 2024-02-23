<template>
  <div>
    <n-qr-code :value="text" />
  </div>
</template>

<script lang="ts" setup>
const text =
  "otpauth://totp/puzzle:defo1215.cn?algorithm=SHA1&digits=6&issuer=puzzle&period=30&secret=K7PA2QG47FCKOT6QNRAXBW4BL67NOTR2";
</script>
