<script lang="ts" setup>
// Cast to any to keep a simple string middleware reference while using TS.
definePageMeta({ middleware: 'auth' })
</script>

<template>
  <div class="p-8">
    <h1 class="text-2xl font-bold mb-4">Dashboard</h1>
  </div>
</template>