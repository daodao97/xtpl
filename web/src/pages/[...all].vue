<route lang="yaml">
meta:
  ssrData: false
</route>

<script setup lang="ts">
import { useRoute } from 'vue-router'

const route = useRoute()
</script>

<template>
  <main>
    <h1>Page not found</h1>
    <p>{{ route.fullPath }}</p>
    <RouterLink to="/">Back home</RouterLink>
  </main>
</template>
