<script lang="ts" setup>
// Cast to any to keep a simple string middleware reference while using TS.
definePageMeta({ middleware: 'auth' })
</script>

<template>
  <div class="p-8 flex gap-8 h-screen">
    <Links></Links>
    <Notes></Notes>
  </div>
</template>